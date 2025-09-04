package database

import "gorm.io/gorm"

func InitRLS(db *gorm.DB) {
	// Enable row level security for the order table
	if err := db.Exec("ALTER TABLE orders ENABLE ROW LEVEL SECURITY").Error; err != nil {
		panic("failed to enable row level security: " + err.Error())
	}

	// Create a policy
	db.Exec(`DROP POLICY IF EXISTS order_rls_policy ON orders`)
	if err := db.Exec(`CREATE POLICY order_rls_policy
	ON orders FOR ALL USING (user_id = current_setting('app.current_user')::int)
	`).Error; err != nil {
		panic("failed to create policy: " + err.Error())
	}
}
