package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/util"
)

func (e *SQLiteLanguage) New() (referenceInicialized interface{}, err error) {
	referenceInicialized = &SQLiteLanguage{}
	err = referenceInicialized.(*SQLiteLanguage).Connect(constants.KSQLiteConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.(*SQLiteLanguage).Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
