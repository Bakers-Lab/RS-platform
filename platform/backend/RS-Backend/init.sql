CREATE TABLE
    IF NOT EXISTS datasets (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        comment TEXT,
        path VARCHAR(255) NOT NULL,
        store_format VARCHAR(50),
        state VARCHAR(50) NOT NULL CHECK (
            state IN ('Failed', 'Ready', 'Deleted')
        )
    );

-- 可以添加更多初始化 SQL 语句
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255)
    );

CREATE TABLE
    IF NOT EXISTS dataset_batches (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        datasetid INT NOT NULL,
        state VARCHAR(50) DEFAULT NULL,
        samples_num INT DEFAULT NULL,
        create_at BIGINT DEFAULT NULL,
        updated_at BIGINT DEFAULT NULL,
        deleted_at BIGINT DEFAULT NULL,
        file_size BIGINT DEFAULT NULL
    );

CREATE TABLE
    IF NOT EXISTS rs_models (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        taskname VARCHAR(255) DEFAULT NULL,
        comment VARCHAR(255) DEFAULT NULL,
        request_url VARCHAR(255) DEFAULT NULL,
        params VARCHAR(255) DEFAULT NULL
    );

CREATE TABLE
    IF NOT EXISTS rs_infer_jobs (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        path VARCHAR(255) DEFAULT NULL,
        state VARCHAR(50) NOT NULL CHECK (
            state IN ('Pending', 'Runing', 'Finished','Failed','Deleted')
        ),
        modelid INT NOT NULL,
        datasetid INT NOT NULL,
        file_size BIGINT DEFAULT NULL,
        create_at BIGINT DEFAULT NULL,
        updated_at BIGINT DEFAULT NULL,
        deleted_at BIGINT DEFAULT NULL
    );

CREATE TABLE
    IF NOT EXISTS rs_eval_jobs (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        path VARCHAR(255) DEFAULT NULL,
        state VARCHAR(50) NOT NULL CHECK (
            state IN ('Pending', 'Runing', 'Finished','Failed','Deleted')
        ),
        inferjobid INT NOT NULL,
        datasetid INT NOT NULL,
        file_size BIGINT DEFAULT NULL,
        create_at BIGINT DEFAULT NULL,
        updated_at BIGINT DEFAULT NULL,
        deleted_at BIGINT DEFAULT NULL
    );