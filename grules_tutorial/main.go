package main

import (
	"fmt"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// a fact is a pointer to a struct instance
type MyFact struct {
	IntAttribute     int64
	StringAttribute  string
	BooleanAttribute bool
	FloatAttribute   float64
	TimeAttribute    time.Time
	WhatToSay        string
}

func (mf *MyFact) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let's say %q \n", sentence)
}

func main() {
	// some insances of MyFact

	myFact1 := &MyFact{
		IntAttribute:     123,
		StringAttribute:  "My string attribute for 123",
		BooleanAttribute: true,
		FloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}

	myFact2 := &MyFact{
		IntAttribute:     456,
		StringAttribute:  "My string attribute for 456",
		BooleanAttribute: false,
		FloatAttribute:   4.567,
		TimeAttribute:    time.Now(),
	}
	// helpful discussion on applying rule to multiple facts:
	// https://github.com/hyperjumptech/grule-rule-engine/discussions/236

	// from docs: DataContext holds all structs instance to be used in rule execution environment.
	dataCtx := ast.NewDataContext()
	err1 := dataCtx.Add("MF1", myFact1)
	err2 := dataCtx.Add("MF2", myFact2)

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	// a KnowledgeBase is a collection of rules sourced from rule defs loaded from multiple sources
	// a KnowledgeLibrary is a collection of KnowledgeBase blueprints
	// RuleBuilder is used to create KnowledgeBase instances and add them to the KnowledgeLibrary

	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	// DSL = Domain-Specific Language
	// basic rule def as a raw string in the DSL
	// for more on the Grule Rule Language (GRL)
	// https://github.com/hyperjumptech/grule-rule-engine/blob/8274ea948544d4e48e001e28e5a01d4e64ff70b0/docs/en/GRL_en.md

	drls := `
	rule CheckValues "Check the default values" salience 10 {
    when 
        MF1.IntAttribute == 123 && MF1.StringAttribute == "My string attribute for 123"
    then
        MF1.WhatToSay = MF1.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
	}
	`
	// Retract makes sure the rule is not immediately re-evaluated (prevents infinite loop?)

	// adding rule def to our knowledgeLibrary from a declared resource
	// naming it 'TutorialRules'  version '0.0.1'

	// convert string to byte slice first, then create a resource from it
	// you can also create resources from files, located at URL endpoints, hosted in git repos, and from JSON
	bs := pkg.NewBytesResource([]byte(drls))
	builderErr := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)

	if builderErr != nil {
		panic(builderErr)
	}

	// getting instance of KnowledgeBase from knowledgeLibrary
	// Each instance you obtain from the knowledgeLibrary is a unique clone from the underlying KnowledgeBase blueprint.
	// see docs for more details on how this an be leveraged
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")

	// execute the KnowledgeBase instance using the prepared DataContext

	engine := engine.NewGruleEngine()

	engineErr := engine.Execute(dataCtx, knowledgeBase)

	if engineErr != nil {
		panic(engineErr)
	}

	fmt.Println(myFact1.WhatToSay)

}
