package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!global`:         `global`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!event`:          `event`,
	`!and`:            `and`,
	`!if`:             `if`,
	`alter`:           `alter`,
	`get`:             `get`,
	`global`:          `global`,
	`swivel-table`:    `swivel-table`,
	`tout`:            `tout`,
	`pt`:              `pt`,
	`f`:               `f`,
	`g`:               `g`,
	`if`:              `if`,
	`catch`:           `catch`,
	`tread`:           `tread`,
	`and`:             `and`,
	`or`:              `or`,
	`try`:             `try`,
	`prepend`:         `prepend`,
	`append`:          `append`,
	`murex-docs`:      `murex-docs`,
	`out`:             `out`,
	`trypipe`:         `trypipe`,
	`brace-quote`:     `brace-quote`,
	`ttyfd`:           `ttyfd`,
	`rx`:              `rx`,
	`export`:          `export`,
	`set`:             `set`,
	`getfile`:         `getfile`,
	`post`:            `post`,
	`err`:             `err`,
	`>`:               `>`,
	`read`:            `read`,
	`event`:           `event`,
	`swivel-datatype`: `swivel-datatype`,
	`>>`:              `>>`,
}
