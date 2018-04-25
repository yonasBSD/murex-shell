package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`unset`:           `unset`,
	`prepend`:         `prepend`,
	`swivel-datatype`: `swivel-datatype`,
	`getfile`:         `getfile`,
	`>>`:              `>>`,
	`f`:               `f`,
	`tread`:           `tread`,
	`swivel-table`:    `swivel-table`,
	`post`:            `post`,
	`out`:             `out`,
	`read`:            `read`,
	`catch`:           `catch`,
	`event`:           `event`,
	`alter`:           `alter`,
	`err`:             `err`,
	`if`:              `if`,
	`set`:             `set`,
	`pt`:              `pt`,
	`or`:              `or`,
	`append`:          `append`,
	`murex-docs`:      `murex-docs`,
	`rx`:              `rx`,
	`g`:               `g`,
	`trypipe`:         `trypipe`,
	`brace-quote`:     `brace-quote`,
	`>`:               `>`,
	`ttyfd`:           `ttyfd`,
	`get`:             `get`,
	`tout`:            `tout`,
	`and`:             `and`,
	`try`:             `try`,
}
