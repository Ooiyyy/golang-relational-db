CREATE TABLE students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    class_id INT,
    FOREIGN KEY (class_id) REFERENCES classes(id)
);