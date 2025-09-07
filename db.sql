DROP DATABASE IF EXISTS reltrace;
CREATE DATABASE reltrace CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE reltrace;

-- =====================================================
-- CORE ENTITIES
-- =====================================================

-- Companies table (root entity)
CREATE TABLE companies (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    industry VARCHAR(100),
    founded_year INT,
    headquarters VARCHAR(255),
    parent_company_id INT NULL, -- Self-reference for subsidiaries
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (parent_company_id) REFERENCES companies(id) ON DELETE SET NULL
);

-- Locations/Offices
CREATE TABLE locations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    city VARCHAR(100),
    country VARCHAR(100),
    is_headquarters BOOLEAN DEFAULT FALSE,
    
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);

-- Departments
CREATE TABLE departments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_id INT NOT NULL,
    location_id INT NULL,
    name VARCHAR(255) NOT NULL,
    budget DECIMAL(15,2),
    parent_department_id INT NULL, -- Self-reference for sub-departments
    
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (location_id) REFERENCES locations(id) ON DELETE SET NULL,
    FOREIGN KEY (parent_department_id) REFERENCES departments(id) ON DELETE SET NULL
);

-- =====================================================
-- PEOPLE & ROLES
-- =====================================================

-- Employees
CREATE TABLE employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_id INT NOT NULL,
    department_id INT NULL,
    location_id INT NULL,
    manager_id INT NULL, -- Self-reference for hierarchy
    employee_number VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(50),
    hire_date DATE NOT NULL,
    salary DECIMAL(10,2),
    status ENUM('active', 'inactive', 'terminated') DEFAULT 'active',
    
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE SET NULL,
    FOREIGN KEY (location_id) REFERENCES locations(id) ON DELETE SET NULL,
    FOREIGN KEY (manager_id) REFERENCES employees(id) ON DELETE SET NULL
);

-- Job Titles/Positions
CREATE TABLE positions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    level ENUM('entry', 'junior', 'mid', 'senior', 'lead', 'manager', 'director', 'executive'),
    min_salary DECIMAL(10,2),
    max_salary DECIMAL(10,2)
);

-- Employee Positions (Many-to-Many with history)
CREATE TABLE employee_positions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    employee_id INT NOT NULL,
    position_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NULL,
    is_current BOOLEAN DEFAULT TRUE,
    
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (position_id) REFERENCES positions(id) ON DELETE CASCADE
);

-- =====================================================
-- PROJECTS & ASSIGNMENTS
-- =====================================================

-- Projects
CREATE TABLE projects (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_id INT NOT NULL,
    department_id INT NULL,
    project_manager_id INT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    budget DECIMAL(15,2),
    status ENUM('planning', 'active', 'on_hold', 'completed', 'cancelled') DEFAULT 'planning',
    parent_project_id INT NULL, -- Sub-projects
    
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE SET NULL,
    FOREIGN KEY (project_manager_id) REFERENCES employees(id) ON DELETE SET NULL,
    FOREIGN KEY (parent_project_id) REFERENCES projects(id) ON DELETE SET NULL
);

-- Project Assignments (Many-to-Many)
CREATE TABLE project_assignments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    project_id INT NOT NULL,
    employee_id INT NOT NULL,
    role VARCHAR(100),
    allocation_percentage INT DEFAULT 100,
    start_date DATE,
    end_date DATE,
    
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);

-- =====================================================
-- CLIENTS & CONTRACTS
-- =====================================================

-- Clients
CREATE TABLE clients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_name VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(50),
    industry VARCHAR(100),
    account_manager_id INT NULL, -- Employee managing this client
    
    FOREIGN KEY (account_manager_id) REFERENCES employees(id) ON DELETE SET NULL
);

-- Contracts
CREATE TABLE contracts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    client_id INT NOT NULL,
    company_id INT NOT NULL,
    project_id INT NULL,
    contract_number VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    value DECIMAL(15,2),
    start_date DATE,
    end_date DATE,
    status ENUM('draft', 'active', 'completed', 'terminated') DEFAULT 'draft',
    
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL
);

-- =====================================================
-- FINANCIAL RECORDS
-- =====================================================

-- Expense Categories
CREATE TABLE expense_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    parent_category_id INT NULL,
    
    FOREIGN KEY (parent_category_id) REFERENCES expense_categories(id) ON DELETE SET NULL
);

-- Expenses
CREATE TABLE expenses (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_id INT NOT NULL,
    employee_id INT NULL,
    project_id INT NULL,
    department_id INT NULL,
    category_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    description TEXT,
    expense_date DATE NOT NULL,
    receipt_url VARCHAR(500),
    status ENUM('pending', 'approved', 'rejected', 'paid') DEFAULT 'pending',
    
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE SET NULL,
    FOREIGN KEY (category_id) REFERENCES expense_categories(id) ON DELETE RESTRICT
);

-- =====================================================
-- AUDIT & LOGGING
-- =====================================================

-- Activity Log (tracks changes to important entities)
CREATE TABLE activity_log (
    id INT PRIMARY KEY AUTO_INCREMENT,
    entity_type VARCHAR(50) NOT NULL,
    entity_id INT NOT NULL,
    user_id INT NULL, -- Which employee made the change
    action ENUM('create', 'update', 'delete') NOT NULL,
    old_values JSON,
    new_values JSON,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (user_id) REFERENCES employees(id) ON DELETE SET NULL
);

-- Insert Companies (with parent-child relationships)
INSERT INTO companies (name, industry, founded_year, headquarters, parent_company_id) VALUES
( 'TechCorp Global', 'Technology', 2010, 'San Francisco, CA', NULL),
( 'TechCorp Europe', 'Technology', 2015, 'London, UK', 1),
( 'TechCorp Asia', 'Technology', 2018, 'Singapore', 1),
( 'DataSoft Inc', 'Software', 2012, 'Austin, TX', NULL),
( 'CloudVision Ltd', 'Cloud Computing', 2019, 'Seattle, WA', NULL);

-- Insert Locations
INSERT INTO locations (company_id, name, address, city, country, is_headquarters) VALUES
(1, 'HQ San Francisco', '123 Tech Street', 'San Francisco', 'USA', TRUE),
(1, 'Austin Office', '456 Innovation Blvd', 'Austin', 'USA', FALSE),
(2, 'London HQ', '789 Digital Ave', 'London', 'UK', TRUE),
(2, 'Berlin Office', '321 Code Street', 'Berlin', 'Germany', FALSE),
(3, 'Singapore HQ', '654 Asia Tech Park', 'Singapore', 'Singapore', TRUE),
(4, 'Austin HQ', '987 Software Lane', 'Austin', 'USA', TRUE),
(5, 'Seattle HQ', '147 Cloud Way', 'Seattle', 'USA', TRUE);

-- Insert Departments (with hierarchy)
INSERT INTO departments (company_id, location_id, name, budget, parent_department_id) VALUES
(1, 1, 'Engineering', 5000000.00, NULL),
(1, 1, 'Backend Development', 2000000.00, 1),
(1, 1, 'Frontend Development', 1500000.00, 1),
(1, 1, 'DevOps', 1000000.00, 1),
(1, 1, 'Marketing', 1000000.00, NULL),
(1, 1, 'Digital Marketing', 600000.00, 5),
(1, 1, 'Sales', 800000.00, NULL),
(2, 3, 'Engineering Europe', 3000000.00, NULL),
(3, 5, 'Engineering Asia', 2000000.00, NULL);

-- Insert Positions
INSERT INTO positions (title, level, min_salary, max_salary) VALUES
('Software Engineer', 'junior', 70000.00, 90000.00),
('Senior Software Engineer', 'senior', 100000.00, 130000.00),
('Lead Software Engineer', 'lead', 130000.00, 160000.00),
('Engineering Manager', 'manager', 150000.00, 200000.00),
('Marketing Specialist', 'mid', 50000.00, 70000.00),
('Sales Representative', 'mid', 45000.00, 65000.00),
('Project Manager', 'manager', 90000.00, 120000.00),
('CEO', 'executive', 300000.00, 500000.00),
('CTO', 'executive', 250000.00, 350000.00),
('DevOps Engineer', 'senior', 95000.00, 125000.00);

-- Insert Employees (with manager hierarchy)
INSERT INTO employees (company_id, department_id, location_id, manager_id, employee_number, first_name, last_name, email, phone, hire_date, salary, status) VALUES
(1, 1, 1, NULL, 'EMP001', 'John', 'Smith', 'john.smith@techcorp.com', '+1-555-0101', '2020-01-15', 180000.00, 'active'),
(1, 2, 1, 1, 'EMP002', 'Sarah', 'Johnson', 'sarah.johnson@techcorp.com', '+1-555-0102', '2020-03-01', 120000.00, 'active'),
(1, 2, 1, 2, 'EMP003', 'Mike', 'Davis', 'mike.davis@techcorp.com', '+1-555-0103', '2021-06-15', 85000.00, 'active'),
(1, 3, 1, 1, 'EMP004', 'Emily', 'Brown', 'emily.brown@techcorp.com', '+1-555-0104', '2021-02-20', 110000.00, 'active'),
(1, 4, 1, 1, 'EMP005', 'David', 'Wilson', 'david.wilson@techcorp.com', '+1-555-0105', '2020-08-10', 115000.00, 'active'),
(1, 5, 1, NULL, 'EMP006', 'Lisa', 'Garcia', 'lisa.garcia@techcorp.com', '+1-555-0106', '2019-11-01', 75000.00, 'active'),
(1, 7, 1, NULL, 'EMP007', 'Robert', 'Miller', 'robert.miller@techcorp.com', '+1-555-0107', '2020-05-12', 65000.00, 'active'),
(2, 8, 3, NULL, 'EMP008', 'Anna', 'Taylor', 'anna.taylor@techcorp.com', '+44-20-1234', '2021-01-10', 95000.00, 'active'),
(3, 9, 5, NULL, 'EMP009', 'James', 'Anderson', 'james.anderson@techcorp.com', '+65-1234-5678', '2021-09-01', 85000.00, 'active'),
(1, NULL, 1, NULL, 'EMP010', 'Alice', 'CEO', 'alice.ceo@techcorp.com', '+1-555-0100', '2010-01-01', 400000.00, 'active');

-- Insert Employee Positions
INSERT INTO employee_positions (employee_id, position_id, start_date, end_date, is_current) VALUES
(1, 4, '2020-01-15', NULL, TRUE),
(2, 3, '2020-03-01', NULL, TRUE),
(3, 1, '2021-06-15', NULL, TRUE),
(4, 2, '2021-02-20', NULL, TRUE),
(5, 10, '2020-08-10', NULL, TRUE),
(6, 5, '2019-11-01', NULL, TRUE),
(7, 6, '2020-05-12', NULL, TRUE),
(8, 2, '2021-01-10', NULL, TRUE),
(9, 1, '2021-09-01', NULL, TRUE),
(10, 8, '2010-01-01', NULL, TRUE);

-- Insert Projects (with parent-child relationships)
INSERT INTO projects (company_id, department_id, project_manager_id, name, description, start_date, end_date, budget, status, parent_project_id) VALUES
(1, 1, 1, 'Platform Redesign', 'Complete platform architecture redesign', '2023-01-01', '2023-12-31', 2000000.00, 'active', NULL),
(1, 2, 2, 'API Development', 'Build new REST API infrastructure', '2023-02-01', '2023-08-31', 800000.00, 'active', 1),
(1, 3, 4, 'UI/UX Overhaul', 'Redesign user interface', '2023-03-01', '2023-10-31', 600000.00, 'active', 1),
(1, 5, 6, 'Marketing Campaign Q2', 'Launch new product marketing', '2023-04-01', '2023-06-30', 300000.00, 'completed', NULL),
(2, 8, 8, 'European Expansion', 'Expand services to European market', '2023-01-15', '2023-11-30', 1500000.00, 'active', NULL);

-- Insert Project Assignments
INSERT INTO project_assignments (project_id, employee_id, role, allocation_percentage, start_date, end_date) VALUES
(1, 1, 'Project Manager', 80, '2023-01-01', NULL),
(1, 2, 'Technical Lead', 100, '2023-01-01', NULL),
(2, 2, 'Lead Developer', 60, '2023-02-01', NULL),
(2, 3, 'Backend Developer', 100, '2023-02-01', NULL),
(2, 5, 'DevOps Engineer', 40, '2023-02-01', NULL),
(3, 4, 'Frontend Lead', 100, '2023-03-01', NULL),
(4, 6, 'Marketing Manager', 100, '2023-04-01', '2023-06-30'),
(5, 8, 'Regional Manager', 100, '2023-01-15', NULL);

-- Insert Clients
INSERT INTO clients (company_name, contact_person, email, phone, industry, account_manager_id) VALUES
('Retail Giant Corp', 'Tom Wilson', 'tom@retailgiant.com', '+1-555-2001', 'Retail', 7),
('Banking Solutions Ltd', 'Maria Rodriguez', 'maria@bankingsol.com', '+1-555-2002', 'Finance', 7),
('Healthcare System Inc', 'Dr. James Lee', 'james@healthsys.com', '+1-555-2003', 'Healthcare', 7),
('Education Platform Co', 'Susan White', 'susan@eduplatform.com', '+1-555-2004', 'Education', 7);

-- Insert Contracts
INSERT INTO contracts (client_id, company_id, project_id, contract_number, title, value, start_date, end_date, status) VALUES
(1, 1, 1, 'CNT-2023-001', 'E-commerce Platform Development', 1500000.00, '2023-01-01', '2023-12-31', 'active'),
(2, 1, 2, 'CNT-2023-002', 'Banking API Integration', 800000.00, '2023-02-01', '2023-08-31', 'active'),
(3, 1, NULL, 'CNT-2023-003', 'Healthcare Management System', 1200000.00, '2023-06-01', '2024-05-31', 'active'),
(4, 2, 5, 'CNT-2023-004', 'Online Learning Platform', 900000.00, '2023-03-01', '2023-11-30', 'active');

-- Insert Expense Categories (with hierarchy)
INSERT INTO expense_categories (name, parent_category_id) VALUES
('Office Expenses', NULL),
('Office Supplies', 1),
('Office Equipment', 1),
('Travel & Entertainment', NULL),
('Business Travel', 4),
('Client Entertainment', 4),
('Professional Services', NULL),
('Legal Services', 7),
('Consulting Services', 7);

-- Insert Expenses
INSERT INTO expenses (company_id, employee_id, project_id, department_id, category_id, amount, description, expense_date, status) VALUES
(1, 3, 2, 2, 2, 150.00, 'Development books and resources', '2023-03-15', 'approved'),
(1, 2, 1, 2, 5, 1200.00, 'Client meeting in New York', '2023-02-20', 'paid'),
(1, 4, 3, 3, 3, 2500.00, 'New design workstation', '2023-03-01', 'approved'),
(1, 1, 1, 1, 8, 5000.00, 'Legal review of contracts', '2023-01-30', 'paid'),
(2, 8, 5, 8, 5, 800.00, 'Travel to Berlin office', '2023-02-10', 'approved');

-- Insert Activity Log
INSERT INTO activity_log (entity_type, entity_id, user_id, action, old_values, new_values) VALUES
('employees', 3, 2, 'update', '{"salary": 80000}', '{"salary": 85000}'),
('projects', 1, 1, 'update', '{"status": "planning"}', '{"status": "active"}'),
('contracts', 1, 7, 'create', NULL, '{"client_id": 1, "value": 1500000}');

-- =====================================================
-- TEST SCENARIOS
-- =====================================================

/*
TEST CASES:

1. SIMPLE HIERARCHY TEST:
   Root: companies, ID: 1 (TechCorp Global)
   Expected: Should pull the parent company + subsidiaries + all related data

2. EMPLOYEE HIERARCHY TEST:
   Root: employees, ID: 1 (John Smith - Engineering Manager)
   Expected: Should pull his reports, projects, expenses, etc.

3. PROJECT DEEP DIVE TEST:
   Root: projects, ID: 1 (Platform Redesign)
   Expected: Should pull project + sub-projects + assignments + contracts + expenses

4. CIRCULAR REFERENCE TEST:
   Root: employees, ID: 2 (Sarah Johnson)
   Expected: Should handle manager->employee relationships without infinite loops

5. MULTI-LEVEL FOREIGN KEY TEST:
   Root: contracts, ID: 1
   Expected: Should traverse client->contract->project->employees->expenses chain

6. SELF-REFERENCE TEST:
   Root: departments, ID: 1 (Engineering)
   Expected: Should pull parent department + all sub-departments + employees
*/