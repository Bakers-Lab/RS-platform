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