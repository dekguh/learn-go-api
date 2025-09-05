CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title VARCHAR(128) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP POLICY IF EXISTS todos_rls_policy ON todos;
ALTER TABLE todos ENABLE ROW LEVEL SECURITY;
ALTER TABLE todos FORCE ROW LEVEL SECURITY;  -- Force row level security for owner table

CREATE POLICY todos_rls_policy ON todos
FOR ALL USING (user_id = current_setting('app.current_user_id')::bigint)
WITH CHECK (user_id = current_setting('app.current_user_id')::bigint);