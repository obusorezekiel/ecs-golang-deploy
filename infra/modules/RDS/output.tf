# Output the DB instance identifier
output "db_instance_identifier" {
  description = "The identifier of the RDS instance"
  value       = aws_db_instance.complex_db.id
}

# Output the database endpoint (hostname)
output "db_instance_endpoint" {
  description = "The connection endpoint for the RDS instance"
  value       = aws_db_instance.complex_db.endpoint
}

