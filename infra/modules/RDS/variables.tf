# Define the amount of storage to allocate for the DB instance
variable "db_storage" {
  description = "The allocated storage for the RDS instance in GB"
  type        = number
}

# Define the database engine version (PostgreSQL in this case)
variable "db_engine_version" {
  description = "The version of the database engine"
  type        = string
}

# Define the instance class for the DB instance
variable "db_instance_class" {
  description = "The instance class for the DB instance"
  type        = string
}

# Define the database name to create
variable "db_name" {
  description = "The name of the database to create"
  type        = string
}

# Define the username for the database
variable "dbuser" {
  description = "The username for the database"
  type        = string
}

# Define the password for the database
variable "dbpassword" {
  description = "The password for the database"
  type        = string
}

# Define the subnet group name where the RDS instance will be launched
variable "db_subnet_group_name" {
  description = "The DB subnet group to associate with the RDS instance"
  type        = string
}

# Define the identifier for the RDS instance
variable "db_identifier" {
  description = "The identifier for the DB instance"
  type        = string
}

variable "skip_db_snapshot" {
  description = "Whether to skip taking a final snapshot when deleting the DB instance"
  type        = bool
}

variable "rds_sg" {
  description = "The security group IDs to assign to the RDS instance"
  type        = list(string)
}
