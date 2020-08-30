// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/sirupsen/logrus"
)

// Debug :
func (object *LogObject) Debug(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Debug(args...)
}

// Print :
func (object *LogObject) Print(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Print(args...)
}

// Info :
func (object *LogObject) Info(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Info(args...)
}

// Warn :
func (object *LogObject) Warn(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Warn(args...)
}

// Warning :
func (object *LogObject) Warning(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Warning(args...)
}

// Error :
func (object *LogObject) Error(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Error(args...)
}

// Panic :
func (object *LogObject) Panic(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Panic(args...)
}

// Fatal :
func (object *LogObject) Fatal(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Fatal(args...)
}

// Debugf :
func (object *LogObject) Debugf(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Debugf(format, args...)
}

// Infof :
func (object *LogObject) Infof(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Infof(format, args...)
}

// Warnf :
func (object *LogObject) Warnf(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Warnf(format, args...)
}

// Warningf :
func (object *LogObject) Warningf(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Warningf(format, args...)
}

// Panicf :
func (object *LogObject) Panicf(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Panicf(format, args...)
}

// Fatalf :
func (object *LogObject) Fatalf(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Fatalf(format, args...)
}

// Errorf :
func (object *LogObject) Errorf(format string, args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Errorf(format, args...)
}

// Debugln :
func (object *LogObject) Debugln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Debugln(args...)
}

// Println :
func (object *LogObject) Println(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Println(args...)
}

// Infoln :
func (object *LogObject) Infoln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Infoln(args...)
}

// Warnln :
func (object *LogObject) Warnln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Warnln(args...)
}

// Warningln :
func (object *LogObject) Warningln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Warningln(args...)
}

// Errorln :
func (object *LogObject) Errorln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Errorln(args...)
}

// Panicln :
func (object *LogObject) Panicln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Panicln(args...)
}

// Fatalln :
func (object *LogObject) Fatalln(args ...interface{}) {
	if !object.Initialized {
		logrus.Fatal("LogObject used without initialization")
		return
	}
	object.logger.WithFields(object.Fields).Fatalln(args...)
}

// Emulate a Trace level - mapped to Warning

// Trace :
func (object *LogObject) Trace(args ...interface{}) {
	if !object.Initialized {
		log.Fatal("LogObject used without initialization")
		return
	}
	log.WithFields(object.Fields).Warn(args...)
}

// Tracef :
func (object *LogObject) Tracef(format string, args ...interface{}) {
	if !object.Initialized {
		log.Fatal("LogObject used without initialization")
		return
	}
	log.WithFields(object.Fields).Warnf(format, args...)
}

// Traceln :
func (object *LogObject) Traceln(args ...interface{}) {
	if !object.Initialized {
		log.Fatal("LogObject used without initialization")
		return
	}
	log.WithFields(object.Fields).Warnln(args...)
}
