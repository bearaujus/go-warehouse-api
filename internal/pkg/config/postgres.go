package config

import (
	"fmt"
)

func (c *Config) generatePostgresDSN(host, port string) string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		host,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDB,
		port,
	)
}

func (c *Config) GetPostgresMasterDSN() string {
	return c.generatePostgresDSN(c.PostgresMasterContainerName, c.PostgresMasterPort)
}

func (c *Config) GetPostgresReplicasDSN() []string {
	type h struct {
		host, port string
	}

	hs := []h{
		{
			host: c.PostgresReplica1ContainerName,
			port: c.PostgresReplica1Port,
		},
		{
			host: c.PostgresReplica2ContainerName,
			port: c.PostgresReplica2Port,
		},
		{
			host: c.PostgresReplica3ContainerName,
			port: c.PostgresReplica3Port,
		},
	}

	ret := make([]string, len(hs))
	for i, d := range hs {
		ret[i] = c.generatePostgresDSN(d.host, d.port)
	}
	return ret
}
