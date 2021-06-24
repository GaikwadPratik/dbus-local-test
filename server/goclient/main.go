package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	go clientHardcodedCall()
	go generatorCall()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Warn().Msg("Adios!")
}

func clientHardcodedCall() error {
	const DEST = "com.hiveio.vmmanager1"
	const DBUS_OBJECT_PATH = dbus.ObjectPath("/com/hiveio/vmmanager1")
	conn, err := dbus.SystemBus()

	if err != nil {
		log.Error().Err(err).Msg("While initiating connection to system bus in initiateServer")
		return err
	}

	demoObject := conn.Object(DEST, DBUS_OBJECT_PATH)

	obj := NewCom_Hiveio_Vm_Manager(demoObject)

	val, err := obj.CheckHostForMigration(context.Background(), "test", "test")

	if err != nil {
		log.Error().Err(err).Msg("While checking host for migration in hardcoded call")
		return err
	}

	log.Debug().Bool("val", val).Msg("Recived response in hardcoded call")

	return nil
}

func generatorCall() error {
	const DEST = "com.hiveio.vmmanager"
	const DBUS_OBJECT_PATH = dbus.ObjectPath("/com/hiveio/vmmanager")
	conn, err := dbus.SystemBus()

	if err != nil {
		log.Error().Err(err).Msg("While initiating connection to system bus in initiateServer")
		return err
	}

	demoObject := conn.Object(DEST, DBUS_OBJECT_PATH)

	obj := NewCom_Hiveio_Vm_Manager(demoObject)

	val, err := obj.CheckHostForMigration(context.Background(), "test", "test")

	if err != nil {
		log.Error().Err(err).Msg("While checking host for migration in generator call")
		return err
	}

	log.Debug().Bool("val", val).Msg("Recived response in generator call")

	return nil
}
