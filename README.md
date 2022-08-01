## Golang Project - Digitalent PROA Batch 3

## Database
```sql
CREATE DATABASE digitalent_task_db;

USE digitalent_task_db;

CREATE TABLE m_task (
	id int primary key auto_increment,
	task varchar(255),
	employee varchar(100),
	deadline date
);
```
### Running Project
```
DB_HOST=localhost DB_PORT=3306 DB_USER=root DB_PASSWORD=P@ssw0rd DB_NAME=digitalent_task_db API_PORT=8888 DB_DRIVER=mysql  go run .
```