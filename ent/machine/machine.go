// Code generated by entc, DO NOT EDIT.

package machine

const (
	// Label holds the string label denoting the machine type in the database.
	Label = "machine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the machine in the database.
	Table = "machines"
)

// Columns holds all SQL columns for machine fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
