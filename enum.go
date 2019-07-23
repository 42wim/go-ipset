package ipset

import (
	"github.com/ti-mo/netfilter"
)

const (
	Protocol = 6
)

type messageType netfilter.MessageType

const (
	_           messageType = iota
	CmdProtocol             //  1: Return protocol version
	CmdCreate               //  2: Create a new (empty) set
	CmdDestroy              //  3: Destroy a (empty) set
	CmdFlush                //  4: Remove all elements from a set
	CmdRename               //  5: Rename a set
	CmdSwap                 //  6: Swap two sets
	CmdList                 //  7: List sets
	CmdSave                 //  8: Save sets
	CmdAdd                  //  9: Add an element to a set
	CmdDel                  // 10: Delete an element from a set
	CmdTest                 // 11: Test an element in a set
	CmdHeader               // 12: Get set header data only
	CmdType                 // 13: Get set type
)

type AttributeType int

const (
	_                  AttributeType = iota
	SetAttrProtocol                  //  1: Protocol version
	SetAttrSetName                   //  2: Name of the set
	SetAttrTypeName                  //  3: Typename
	SetAttrRevision                  //  4: Settype revision
	SetAttrFamily                    //  5: Settype family
	SetAttrFlags                     //  6: Flags at command level
	SetAttrData                      //  7: Nested attributes
	SetAttrADT                       //  8: Multiple data containers
	SetAttrLineNo                    //  9: Restore lineno
	SetAttrProtocolMin               // 10: Minimal supported version number
	SetAttrMax

	SetAttrSetName2    = SetAttrTypeName    // Setname at rename/swap
	SetAttrRevisionMin = SetAttrProtocolMin // type rev min
)
