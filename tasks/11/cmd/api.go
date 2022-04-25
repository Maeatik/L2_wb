package cmd

import (
	"L2/tasks/11/api"
	"L2/tasks/11/domain"
	"github.com/spf13/cobra"
	"time"
)

var (
	pHost           string
	pPort           string
	pReadTimeout    uint8
	pWriteTimeout   uint8
	pMaxHeaderBytes int

	//Настраиваем сервер
	apiServerCmd = &cobra.Command{
		Use:   "api-server",
		Short: "Run REST API server",
		Run: func(cmd *cobra.Command, args []string) {
			config := &api.ServerConfig{
				Host:           pHost,
				Port:           pPort,
				ReadTimeout:    time.Duration(pReadTimeout) * time.Second,
				WriteTimeout:   time.Duration(pWriteTimeout) * time.Second,
				MaxHeaderBytes: pMaxHeaderBytes,
			}
			//создаем хранилище моделей
			storage := domain.NewStorage()
			//стартуем
			api.StartServer(storage, config)
		},
	}
)

//Функция, определяющаяся при инициализации
func init() {
	apiServerCmd.PersistentFlags().StringVar(&pHost, "host", "", "Listening host")
	apiServerCmd.PersistentFlags().StringVar(&pPort, "port", "8000", "Listening port")
	apiServerCmd.PersistentFlags().Uint8Var(&pReadTimeout, "read-timeout", 10, "Read timeout in sec")
	apiServerCmd.PersistentFlags().Uint8Var(&pWriteTimeout, "write-timeout", 10, "Write timeout in sec")
	apiServerCmd.PersistentFlags().IntVar(&pMaxHeaderBytes, "max-header-size", 1<<20, "Max header size in bytes")
}
