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
	go clientManualCall()
	go generatorCall()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Warn().Msg("Adios!")
}

func clientManualCall() error {
	DEST := "com.hiveio.vmmanagerhardcoded"
	DBUS_OBJECT_PATH := dbus.ObjectPath("/com/hiveio/vmmanagerhardcoded")
	conn, err := dbus.SystemBus()

	if err != nil {
		log.Error().Err(err).Msg("While initiating connection to system bus in initiateServer")
		return err
	}

	demoObject := conn.Object(DEST, DBUS_OBJECT_PATH)

	var val bool
	err = demoObject.CallWithContext(context.Background(), "CheckHostForMigration", 0, "test", "test").Store(&val)
	if err != nil {
		log.Error().Err(err).Msg("While checking host for migration in hardcoded call")
		return err
	}
	log.Debug().Bool("val", val).Msg("Recived response in hardcoded manual call")

	DEST = "com.hiveio.vmmanager"
	DBUS_OBJECT_PATH = dbus.ObjectPath("/com/hiveio/vmmanager")
	demoObject = conn.Object(DEST, DBUS_OBJECT_PATH)
	err = demoObject.CallWithContext(context.Background(), "CheckHostForMigration", 0, "test", "test").Store(&val)
	if err != nil {
		log.Error().Err(err).Msg("While checking host for migration in generated call")
		return err
	}
	log.Debug().Bool("val", val).Msg("Recived response in generator manual call")
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

	obj := NewComHiveioVmmanager(demoObject)

	val, err := obj.CheckHostForMigration(context.Background(), "test", "test")

	if err != nil {
		log.Error().Err(err).Msg("While checking host for migration in generator call")
		return err
	}

	log.Debug().Bool("val", val).Msg("Recived response in generator call")

	return nil
}
