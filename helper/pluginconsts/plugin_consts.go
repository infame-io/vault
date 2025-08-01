// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package pluginconsts

// These consts live outside the plugin registry files to prevent import cycles.
const (
	AuthTypeAliCloud         = "alicloud"
	AuthTypeAppId            = "app-id"
	AuthTypeAWS              = "aws"
	AuthTypeAzure            = "azure"
	AuthTypeCF               = "cf"
	AuthTypeGCP              = "gcp"
	AuthTypeGitHub           = "github"
	AuthTypeKerberos         = "kerberos"
	AuthTypeKubernetes       = "kubernetes"
	AuthTypeLDAP             = "ldap"
	AuthTypeOCI              = "oci"
	AuthTypeOkta             = "okta"
	AuthTypePCF              = "pcf"
	AuthTypeRadius           = "radius"
	AuthTypeToken            = "token"
	AuthTypeCert             = "cert"
	AuthTypeOIDC             = "oidc"
	AuthTypeUserpass         = "userpass"
	AuthTypeSAML             = "saml"
	AuthTypeApprole          = "approle"
	AuthTypeJWT              = "jwt"
	SecretEngineAD           = "ad"
	SecretEngineAlicloud     = "alicloud"
	SecretEngineAWS          = "aws"
	SecretEngineAzure        = "azure"
	SecretEngineCassandra    = "cassandra"
	SecretEngineConsul       = "consul"
	SecretEngineGCP          = "gcp"
	SecretEngineGCPKMS       = "gcpkms"
	SecretEngineKubernetes   = "kubernetes"
	SecretEngineMongoDB      = "mongodb"
	SecretEngineMongoDBAtlas = "mongodbatlas"
	SecretEngineMSSQL        = "mssql"
	SecretEngineMySQL        = "mysql"
	SecretEngineNomad        = "nomad"
	SecretEngineOpenLDAP     = "openldap"
	SecretEngineLDAP         = "ldap"
	SecretEnginePostgresql   = "postgresql"
	SecretEngineRabbitMQ     = "rabbitmq"
	SecretEngineTerraform    = "terraform"
	SecretEngineTOTP         = "totp"
	SecretEngineKV           = "kv"
	SecretEngineTransform    = "transform"
	SecretEngineKMIP         = "kmip"
	SecretEngineKeymgmt      = "keymgmt"
	SecretEnginePki          = "pki"
	SecretEngineTransit      = "transit"
	SecretEngineSsh          = "ssh"
	SecretEngineCubbyhole    = "cubbyhole"
	SecretEngineIdentity     = "identity"
	SecretEngineSystem       = "system"
	// SecretEngineGeneric is a very old and deprecated version of KV, but is left
	// for completeness.
	SecretEngineGeneric = "generic"
	// SecretEngineDatabase is the entry type for all databases, i.e. this is the combined
	// database type for every database.
	SecretEngineDatabase = "database"
)

// These DB consts match the type returned from database plugin's Type() method.
const (
	DbCassandraPluginType        = "cassandra"
	DbCouchbasePluginType        = "couchbase"
	DbElasticSearchPluginType    = "elasticsearch"
	DbHanaPluginType             = "hdb"
	DbInfluxDBPluginType         = "influxdb"
	DbMongoDBPluginType          = "mongodb"
	DbMongoAtlasPluginType       = "mongodbatlas"
	DbMsSQLPluginType            = "mssql"
	DbMySQLPluginType            = "mysql"
	DbPostgresqlPluginType       = "pgx"
	DbRedshiftPluginType         = "redshift"
	DbRedisPluginType            = "redis"
	DbRedisElasticachePluginType = "redisElastiCache"
	DbSnowflakePluginType        = "snowflake"
)
