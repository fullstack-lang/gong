package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__ActorState__00000000_ := (&models.ActorState{Name: `A1`}).Stage(stage)
	__ActorState__00000001_ := (&models.ActorState{Name: `A2`}).Stage(stage)
	__ActorState__00000002_ := (&models.ActorState{Name: `Political Personality of Soviet Power is amenable to liberal economy`}).Stage(stage)
	__ActorState__00000003_ := (&models.ActorState{Name: `Political Personality of Soviet Power is dictatorial`}).Stage(stage)
	__ActorState__00000004_ := (&models.ActorState{Name: `Political Personality of Soviet Power is paranoïd`}).Stage(stage)
	__ActorState__00000005_ := (&models.ActorState{Name: `Political Personality of Soviet is nebulous, visionary and impractical`}).Stage(stage)
	__ActorState__00000006_ := (&models.ActorState{Name: `Rick is cynical`}).Stage(stage)
	__ActorState__00000007_ := (&models.ActorState{Name: `Rick is empathetic`}).Stage(stage)
	__ActorState__00000008_ := (&models.ActorState{Name: `Student 0%`}).Stage(stage)
	__ActorState__00000009_ := (&models.ActorState{Name: `Student 5%`}).Stage(stage)

	__ActorStateShape__00000004_ := (&models.ActorStateShape{Name: `A1`}).Stage(stage)
	__ActorStateShape__00000005_ := (&models.ActorStateShape{Name: `A2`}).Stage(stage)
	__ActorStateShape__00000007_ := (&models.ActorStateShape{Name: `Political Personality of Soviet Power is dictatorial`}).Stage(stage)
	__ActorStateShape__00000008_ := (&models.ActorStateShape{Name: `Political Personality of Soviet Power is paranoïd`}).Stage(stage)
	__ActorStateShape__00000009_ := (&models.ActorStateShape{Name: `Political Personality of Soviet is nebulous`}).Stage(stage)
	__ActorStateShape__00000010_ := (&models.ActorStateShape{Name: `PoliticalPersonality ofSoviet Power is amenable to liberal economy`}).Stage(stage)
	__ActorStateShape__00000011_ := (&models.ActorStateShape{Name: `Rick is cynical`}).Stage(stage)
	__ActorStateShape__00000012_ := (&models.ActorStateShape{Name: `Rick is cynical`}).Stage(stage)
	__ActorStateShape__00000013_ := (&models.ActorStateShape{Name: `Rick is empathic`}).Stage(stage)
	__ActorStateShape__00000014_ := (&models.ActorStateShape{Name: `Rick is empathic`}).Stage(stage)
	__ActorStateShape__00000015_ := (&models.ActorStateShape{Name: `Student 0%`}).Stage(stage)
	__ActorStateShape__00000016_ := (&models.ActorStateShape{Name: `Student 5%`}).Stage(stage)

	__ActorStateTransition__00000000_ := (&models.ActorStateTransition{Name: ``}).Stage(stage)
	__ActorStateTransition__00000001_ := (&models.ActorStateTransition{Name: `A1 to A2`}).Stage(stage)
	__ActorStateTransition__00000002_ := (&models.ActorStateTransition{Name: `A1 to A2`}).Stage(stage)
	__ActorStateTransition__00000003_ := (&models.ActorStateTransition{Name: `Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy`}).Stage(stage)
	__ActorStateTransition__00000004_ := (&models.ActorStateTransition{Name: `PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd`}).Stage(stage)
	__ActorStateTransition__00000005_ := (&models.ActorStateTransition{Name: `Rick is cynical to Rick is empathic`}).Stage(stage)
	__ActorStateTransition__00000006_ := (&models.ActorStateTransition{Name: `Student 0% to Student 5%`}).Stage(stage)
	__ActorStateTransition__00000007_ := (&models.ActorStateTransition{Name: `Student 5% to Student 15%`}).Stage(stage)

	__ActorStateTransitionShape__00000000_ := (&models.ActorStateTransitionShape{Name: ``}).Stage(stage)
	__ActorStateTransitionShape__00000001_ := (&models.ActorStateTransitionShape{Name: `A1 to A2`}).Stage(stage)
	__ActorStateTransitionShape__00000002_ := (&models.ActorStateTransitionShape{Name: `Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy`}).Stage(stage)
	__ActorStateTransitionShape__00000003_ := (&models.ActorStateTransitionShape{Name: `PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd`}).Stage(stage)
	__ActorStateTransitionShape__00000004_ := (&models.ActorStateTransitionShape{Name: `Rick is cynical to Rick is empathic`}).Stage(stage)
	__ActorStateTransitionShape__00000005_ := (&models.ActorStateTransitionShape{Name: `Rick is cynical to Rick is empathic`}).Stage(stage)
	__ActorStateTransitionShape__00000006_ := (&models.ActorStateTransitionShape{Name: `Student 0% to Student 5%`}).Stage(stage)

	__Analysis__00000000_ := (&models.Analysis{Name: `Casablanca`}).Stage(stage)
	__Analysis__00000001_ := (&models.Analysis{Name: `Horizon(s) analysis`}).Stage(stage)
	__Analysis__00000002_ := (&models.Analysis{Name: `Kennan's Long Telegram`}).Stage(stage)
	__Analysis__00000003_ := (&models.Analysis{Name: `US Elections`}).Stage(stage)

	__ControlPointShape__00000000_ := (&models.ControlPointShape{Name: `Control Point Shape in  0`}).Stage(stage)
	__ControlPointShape__00000001_ := (&models.ControlPointShape{Name: `Control Point Shape in  0`}).Stage(stage)
	__ControlPointShape__00000002_ := (&models.ControlPointShape{Name: `Control Point Shape in  1`}).Stage(stage)
	__ControlPointShape__00000003_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 0`}).Stage(stage)
	__ControlPointShape__00000004_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 1`}).Stage(stage)
	__ControlPointShape__00000005_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 2`}).Stage(stage)
	__ControlPointShape__00000006_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 3`}).Stage(stage)
	__ControlPointShape__00000007_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 3`}).Stage(stage)
	__ControlPointShape__00000008_ := (&models.ControlPointShape{Name: `Control Point Shape in  2`}).Stage(stage)
	__ControlPointShape__00000009_ := (&models.ControlPointShape{Name: `Control Point Shape in  2`}).Stage(stage)
	__ControlPointShape__00000010_ := (&models.ControlPointShape{Name: `Control Point Shape in  2`}).Stage(stage)
	__ControlPointShape__00000011_ := (&models.ControlPointShape{Name: `Control Point Shape in  0`}).Stage(stage)
	__ControlPointShape__00000012_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 1`}).Stage(stage)
	__ControlPointShape__00000013_ := (&models.ControlPointShape{Name: `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 2`}).Stage(stage)
	__ControlPointShape__00000014_ := (&models.ControlPointShape{Name: `Control Point Shape in  1`}).Stage(stage)
	__ControlPointShape__00000015_ := (&models.ControlPointShape{Name: `Control Point Shape in  1`}).Stage(stage)
	__ControlPointShape__00000016_ := (&models.ControlPointShape{Name: `Control Point Shape in  2`}).Stage(stage)
	__ControlPointShape__00000017_ := (&models.ControlPointShape{Name: `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 0`}).Stage(stage)
	__ControlPointShape__00000018_ := (&models.ControlPointShape{Name: `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 1`}).Stage(stage)
	__ControlPointShape__00000019_ := (&models.ControlPointShape{Name: `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 2`}).Stage(stage)
	__ControlPointShape__00000020_ := (&models.ControlPointShape{Name: `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 3`}).Stage(stage)
	__ControlPointShape__00000021_ := (&models.ControlPointShape{Name: `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 4`}).Stage(stage)
	__ControlPointShape__00000022_ := (&models.ControlPointShape{Name: `Control Point Shape in Student 0% to Student 5% 0`}).Stage(stage)
	__ControlPointShape__00000023_ := (&models.ControlPointShape{Name: `Control Point Shape in Student 0% to Student 5% 1`}).Stage(stage)
	__ControlPointShape__00000024_ := (&models.ControlPointShape{Name: `Control Point Shape in Student 0% to Student 5% 2`}).Stage(stage)
	__ControlPointShape__00000025_ := (&models.ControlPointShape{Name: `Control Point Shape in Student 0% to Student 5% 3`}).Stage(stage)
	__ControlPointShape__00000026_ := (&models.ControlPointShape{Name: `Control Point Shape in Student 0% to Student 5% 4`}).Stage(stage)
	__ControlPointShape__00000027_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 0`}).Stage(stage)
	__ControlPointShape__00000028_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 1`}).Stage(stage)
	__ControlPointShape__00000029_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 2`}).Stage(stage)
	__ControlPointShape__00000030_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 3`}).Stage(stage)
	__ControlPointShape__00000031_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 4`}).Stage(stage)
	__ControlPointShape__00000032_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 0`}).Stage(stage)
	__ControlPointShape__00000033_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 1`}).Stage(stage)
	__ControlPointShape__00000034_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 2`}).Stage(stage)
	__ControlPointShape__00000035_ := (&models.ControlPointShape{Name: `Control Point Shape in Rick is cynical to Rick is empathic 3`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Ilsa's trajectory`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `New Diagram`}).Stage(stage)
	__Diagram__00000002_ := (&models.Diagram{Name: `New Diagram`}).Stage(stage)
	__Diagram__00000003_ := (&models.Diagram{Name: `Reference penetration forcast`}).Stage(stage)
	__Diagram__00000004_ := (&models.Diagram{Name: `Rick's Trajectory`}).Stage(stage)

	__EvolutionDirection__00000000_ := (&models.EvolutionDirection{Name: `Bottom`}).Stage(stage)
	__EvolutionDirection__00000001_ := (&models.EvolutionDirection{Name: `Cynical`}).Stage(stage)
	__EvolutionDirection__00000002_ := (&models.EvolutionDirection{Name: `Democratic`}).Stage(stage)
	__EvolutionDirection__00000003_ := (&models.EvolutionDirection{Name: `Dictactorship`}).Stage(stage)
	__EvolutionDirection__00000004_ := (&models.EvolutionDirection{Name: `Empathetic`}).Stage(stage)
	__EvolutionDirection__00000005_ := (&models.EvolutionDirection{Name: `High Penetration`}).Stage(stage)
	__EvolutionDirection__00000006_ := (&models.EvolutionDirection{Name: `Low Penetration`}).Stage(stage)
	__EvolutionDirection__00000007_ := (&models.EvolutionDirection{Name: `Top`}).Stage(stage)

	__EvolutionDirectionShape__00000000_ := (&models.EvolutionDirectionShape{Name: `Bottom`}).Stage(stage)
	__EvolutionDirectionShape__00000001_ := (&models.EvolutionDirectionShape{Name: `Cynical`}).Stage(stage)
	__EvolutionDirectionShape__00000002_ := (&models.EvolutionDirectionShape{Name: `Cynical`}).Stage(stage)
	__EvolutionDirectionShape__00000004_ := (&models.EvolutionDirectionShape{Name: `Democratic`}).Stage(stage)
	__EvolutionDirectionShape__00000005_ := (&models.EvolutionDirectionShape{Name: `Dictactorship`}).Stage(stage)
	__EvolutionDirectionShape__00000006_ := (&models.EvolutionDirectionShape{Name: `Emphatic`}).Stage(stage)
	__EvolutionDirectionShape__00000007_ := (&models.EvolutionDirectionShape{Name: `Emphatic`}).Stage(stage)
	__EvolutionDirectionShape__00000008_ := (&models.EvolutionDirectionShape{Name: `High Penetration`}).Stage(stage)
	__EvolutionDirectionShape__00000009_ := (&models.EvolutionDirectionShape{Name: `Low Penetration`}).Stage(stage)
	__EvolutionDirectionShape__00000010_ := (&models.EvolutionDirectionShape{Name: `Top`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Parameter__00000000_ := (&models.Parameter{Name: `Abroad capitalistic menace`}).Stage(stage)
	__Parameter__00000001_ := (&models.Parameter{Name: `Anti-tsarist rationalization`}).Stage(stage)
	__Parameter__00000002_ := (&models.Parameter{Name: `Blogs`}).Stage(stage)
	__Parameter__00000003_ := (&models.Parameter{Name: `Capitalism --> Imperialism -> War `}).Stage(stage)
	__Parameter__00000004_ := (&models.Parameter{Name: `Communists are a tiny portion of the population`}).Stage(stage)
	__Parameter__00000005_ := (&models.Parameter{Name: `Conservatisme`}).Stage(stage)
	__Parameter__00000006_ := (&models.Parameter{Name: `Ilsa is back in his life ...`}).Stage(stage)
	__Parameter__00000007_ := (&models.Parameter{Name: `Interventions`}).Stage(stage)
	__Parameter__00000008_ := (&models.Parameter{Name: `Lenin's Need to protect proletarians from all countries`}).Stage(stage)
	__Parameter__00000009_ := (&models.Parameter{Name: `Mysteriously gloomy`}).Stage(stage)
	__Parameter__00000010_ := (&models.Parameter{Name: `P1`}).Stage(stage)
	__Parameter__00000011_ := (&models.Parameter{Name: `P2`}).Stage(stage)
	__Parameter__00000012_ := (&models.Parameter{Name: `Paranoïd view of foreign powers (self sustained)`}).Stage(stage)
	__Parameter__00000013_ := (&models.Parameter{Name: `Publications`}).Stage(stage)
	__Parameter__00000014_ := (&models.Parameter{Name: `Rejection of early war communism`}).Stage(stage)
	__Parameter__00000015_ := (&models.Parameter{Name: `Russian Asiatic attraction to imperialism`}).Stage(stage)
	__Parameter__00000016_ := (&models.Parameter{Name: `Siloviki's incitation to maintain internal oppression`}).Stage(stage)
	__Parameter__00000017_ := (&models.Parameter{Name: `Stalin's intolerance to power sharing`}).Stage(stage)

	__ParameterShape__00000000_ := (&models.ParameterShape{Name: `Blogs`}).Stage(stage)
	__ParameterShape__00000001_ := (&models.ParameterShape{Name: `Capitalism --> Imperialism -> War `}).Stage(stage)
	__ParameterShape__00000002_ := (&models.ParameterShape{Name: `Circumstances`}).Stage(stage)
	__ParameterShape__00000003_ := (&models.ParameterShape{Name: `Communists are a tiny portion of the population`}).Stage(stage)
	__ParameterShape__00000004_ := (&models.ParameterShape{Name: `Conservatisme`}).Stage(stage)
	__ParameterShape__00000006_ := (&models.ParameterShape{Name: `Foreign capitalistic menace`}).Stage(stage)
	__ParameterShape__00000007_ := (&models.ParameterShape{Name: `Ilsa is back in his life`}).Stage(stage)
	__ParameterShape__00000008_ := (&models.ParameterShape{Name: `Ilsa is bak in his life`}).Stage(stage)
	__ParameterShape__00000009_ := (&models.ParameterShape{Name: `Interventions`}).Stage(stage)
	__ParameterShape__00000010_ := (&models.ParameterShape{Name: `Lenin's Need to protect proletarians from all countries`}).Stage(stage)
	__ParameterShape__00000012_ := (&models.ParameterShape{Name: `Mysteriously gloomy`}).Stage(stage)
	__ParameterShape__00000013_ := (&models.ParameterShape{Name: `P1`}).Stage(stage)
	__ParameterShape__00000014_ := (&models.ParameterShape{Name: `P2`}).Stage(stage)
	__ParameterShape__00000015_ := (&models.ParameterShape{Name: `Paranoïd view of foreign powers (self sustained)`}).Stage(stage)
	__ParameterShape__00000016_ := (&models.ParameterShape{Name: `Publications`}).Stage(stage)
	__ParameterShape__00000017_ := (&models.ParameterShape{Name: `Rejection of early war communism`}).Stage(stage)
	__ParameterShape__00000018_ := (&models.ParameterShape{Name: `Russian Asiatic attraction to imperialism`}).Stage(stage)
	__ParameterShape__00000019_ := (&models.ParameterShape{Name: `Siloviki's incitation to maintain internal oppression`}).Stage(stage)
	__ParameterShape__00000020_ := (&models.ParameterShape{Name: `Stalin's intolerance to power sharing`}).Stage(stage)

	__ParametersAggregate__00000000_ := (&models.ParametersAggregate{Name: ``}).Stage(stage)
	__ParametersAggregate__00000001_ := (&models.ParametersAggregate{Name: `Circumstances`}).Stage(stage)
	__ParametersAggregate__00000002_ := (&models.ParametersAggregate{Name: `Evangelisation`}).Stage(stage)
	__ParametersAggregate__00000003_ := (&models.ParametersAggregate{Name: `Marxian ideology`}).Stage(stage)
	__ParametersAggregate__00000004_ := (&models.ParametersAggregate{Name: `P1+P2`}).Stage(stage)

	__ParametersAggregateShape__00000000_ := (&models.ParametersAggregateShape{Name: ``}).Stage(stage)
	__ParametersAggregateShape__00000001_ := (&models.ParametersAggregateShape{Name: `Circumstances`}).Stage(stage)
	__ParametersAggregateShape__00000003_ := (&models.ParametersAggregateShape{Name: `Evangelisation`}).Stage(stage)
	__ParametersAggregateShape__00000004_ := (&models.ParametersAggregateShape{Name: `Marxian ideology`}).Stage(stage)

	__Scenario__00000000_ := (&models.Scenario{Name: `Kennan's analysis`}).Stage(stage)
	__Scenario__00000001_ := (&models.Scenario{Name: `Nicky Halley trajectory`}).Stage(stage)
	__Scenario__00000002_ := (&models.Scenario{Name: `Penetration Forcast`}).Stage(stage)
	__Scenario__00000003_ := (&models.Scenario{Name: `Scenario avec intervention sociale concertée`}).Stage(stage)
	__Scenario__00000004_ := (&models.Scenario{Name: `Scénario d'évolution multi domaines sans intervention`}).Stage(stage)
	__Scenario__00000005_ := (&models.Scenario{Name: `The Movie`}).Stage(stage)

	__Workspace__00000000_ := (&models.Workspace{Name: `Default`}).Stage(stage)

	// insertion point for initialization of values

	__ActorState__00000000_.Name = `A1`
	__ActorState__00000000_.Description = ``
	__ActorState__00000000_.IsWithProbaility = false
	__ActorState__00000000_.Probability = ""
	__ActorState__00000000_.ComputedPrefix = ``
	__ActorState__00000000_.IsExpanded = false

	__ActorState__00000001_.Name = `A2`
	__ActorState__00000001_.Description = ``
	__ActorState__00000001_.IsWithProbaility = false
	__ActorState__00000001_.Probability = ""
	__ActorState__00000001_.ComputedPrefix = ``
	__ActorState__00000001_.IsExpanded = false

	__ActorState__00000002_.Name = `Political Personality of Soviet Power is amenable to liberal economy`
	__ActorState__00000002_.Description = ``
	__ActorState__00000002_.IsWithProbaility = false
	__ActorState__00000002_.Probability = ""
	__ActorState__00000002_.ComputedPrefix = ``
	__ActorState__00000002_.IsExpanded = false

	__ActorState__00000003_.Name = `Political Personality of Soviet Power is dictatorial`
	__ActorState__00000003_.Description = ``
	__ActorState__00000003_.IsWithProbaility = false
	__ActorState__00000003_.Probability = ""
	__ActorState__00000003_.ComputedPrefix = ``
	__ActorState__00000003_.IsExpanded = false

	__ActorState__00000004_.Name = `Political Personality of Soviet Power is paranoïd`
	__ActorState__00000004_.Description = ``
	__ActorState__00000004_.IsWithProbaility = false
	__ActorState__00000004_.Probability = ""
	__ActorState__00000004_.ComputedPrefix = ``
	__ActorState__00000004_.IsExpanded = false

	__ActorState__00000005_.Name = `Political Personality of Soviet is nebulous, visionary and impractical`
	__ActorState__00000005_.Description = ``
	__ActorState__00000005_.IsWithProbaility = false
	__ActorState__00000005_.Probability = ""
	__ActorState__00000005_.ComputedPrefix = ``
	__ActorState__00000005_.IsExpanded = false

	__ActorState__00000006_.Name = `Rick is cynical`
	__ActorState__00000006_.Description = ``
	__ActorState__00000006_.IsWithProbaility = true
	__ActorState__00000006_.Probability = models.PERCENT_100
	__ActorState__00000006_.ComputedPrefix = ``
	__ActorState__00000006_.IsExpanded = false

	__ActorState__00000007_.Name = `Rick is empathetic`
	__ActorState__00000007_.Description = ``
	__ActorState__00000007_.IsWithProbaility = true
	__ActorState__00000007_.Probability = models.PERCENT_75
	__ActorState__00000007_.ComputedPrefix = ``
	__ActorState__00000007_.IsExpanded = false

	__ActorState__00000008_.Name = `Student 0%`
	__ActorState__00000008_.Description = ``
	__ActorState__00000008_.IsWithProbaility = true
	__ActorState__00000008_.Probability = models.PERCENT_100
	__ActorState__00000008_.ComputedPrefix = ``
	__ActorState__00000008_.IsExpanded = false

	__ActorState__00000009_.Name = `Student 5%`
	__ActorState__00000009_.Description = ``
	__ActorState__00000009_.IsWithProbaility = true
	__ActorState__00000009_.Probability = models.PERCENT_75
	__ActorState__00000009_.ComputedPrefix = ``
	__ActorState__00000009_.IsExpanded = false

	__ActorStateShape__00000004_.Name = `A1`
	__ActorStateShape__00000004_.X = 96.000000
	__ActorStateShape__00000004_.Y = 369.000000
	__ActorStateShape__00000004_.Width = 240.000000
	__ActorStateShape__00000004_.Height = 80.000000
	__ActorStateShape__00000004_.IsHidden = false

	__ActorStateShape__00000005_.Name = `A2`
	__ActorStateShape__00000005_.X = 776.000000
	__ActorStateShape__00000005_.Y = 194.000000
	__ActorStateShape__00000005_.Width = 240.000000
	__ActorStateShape__00000005_.Height = 80.000000
	__ActorStateShape__00000005_.IsHidden = false

	__ActorStateShape__00000007_.Name = `Political Personality of Soviet Power is dictatorial`
	__ActorStateShape__00000007_.X = 208.000000
	__ActorStateShape__00000007_.Y = 128.000015
	__ActorStateShape__00000007_.Width = 189.000000
	__ActorStateShape__00000007_.Height = 80.000000
	__ActorStateShape__00000007_.IsHidden = false

	__ActorStateShape__00000008_.Name = `Political Personality of Soviet Power is paranoïd`
	__ActorStateShape__00000008_.X = 951.000000
	__ActorStateShape__00000008_.Y = 53.000023
	__ActorStateShape__00000008_.Width = 240.000000
	__ActorStateShape__00000008_.Height = 80.000000
	__ActorStateShape__00000008_.IsHidden = false

	__ActorStateShape__00000009_.Name = `Political Personality of Soviet is nebulous`
	__ActorStateShape__00000009_.X = 19.000000
	__ActorStateShape__00000009_.Y = 276.000000
	__ActorStateShape__00000009_.Width = 167.000000
	__ActorStateShape__00000009_.Height = 109.000000
	__ActorStateShape__00000009_.IsHidden = false

	__ActorStateShape__00000010_.Name = `PoliticalPersonality ofSoviet Power is amenable to liberal economy`
	__ActorStateShape__00000010_.X = 460.000000
	__ActorStateShape__00000010_.Y = 250.000000
	__ActorStateShape__00000010_.Width = 240.000000
	__ActorStateShape__00000010_.Height = 64.000000
	__ActorStateShape__00000010_.IsHidden = false

	__ActorStateShape__00000011_.Name = `Rick is cynical`
	__ActorStateShape__00000011_.X = 87.000000
	__ActorStateShape__00000011_.Y = 545.000000
	__ActorStateShape__00000011_.Width = 198.000000
	__ActorStateShape__00000011_.Height = 74.400000
	__ActorStateShape__00000011_.IsHidden = false

	__ActorStateShape__00000012_.Name = `Rick is cynical`
	__ActorStateShape__00000012_.X = 97.000000
	__ActorStateShape__00000012_.Y = 389.000000
	__ActorStateShape__00000012_.Width = 240.000000
	__ActorStateShape__00000012_.Height = 80.000000
	__ActorStateShape__00000012_.IsHidden = false

	__ActorStateShape__00000013_.Name = `Rick is empathic`
	__ActorStateShape__00000013_.X = 664.000000
	__ActorStateShape__00000013_.Y = 90.000000
	__ActorStateShape__00000013_.Width = 190.901235
	__ActorStateShape__00000013_.Height = 74.643750
	__ActorStateShape__00000013_.IsHidden = false

	__ActorStateShape__00000014_.Name = `Rick is empathic`
	__ActorStateShape__00000014_.X = 643.000000
	__ActorStateShape__00000014_.Y = 210.999985
	__ActorStateShape__00000014_.Width = 254.000000
	__ActorStateShape__00000014_.Height = 68.000000
	__ActorStateShape__00000014_.IsHidden = false

	__ActorStateShape__00000015_.Name = `Student 0%`
	__ActorStateShape__00000015_.X = 45.000000
	__ActorStateShape__00000015_.Y = 283.000000
	__ActorStateShape__00000015_.Width = 184.962963
	__ActorStateShape__00000015_.Height = 76.989189
	__ActorStateShape__00000015_.IsHidden = false

	__ActorStateShape__00000016_.Name = `Student 5%`
	__ActorStateShape__00000016_.X = 818.000000
	__ActorStateShape__00000016_.Y = 58.000000
	__ActorStateShape__00000016_.Width = 159.616667
	__ActorStateShape__00000016_.Height = 80.000000
	__ActorStateShape__00000016_.IsHidden = false

	__ActorStateTransition__00000000_.Name = ``
	__ActorStateTransition__00000000_.ComputedPrefix = ``
	__ActorStateTransition__00000000_.IsExpanded = false

	__ActorStateTransition__00000001_.Name = `A1 to A2`
	__ActorStateTransition__00000001_.ComputedPrefix = ``
	__ActorStateTransition__00000001_.IsExpanded = false

	__ActorStateTransition__00000002_.Name = `A1 to A2`
	__ActorStateTransition__00000002_.ComputedPrefix = ``
	__ActorStateTransition__00000002_.IsExpanded = false

	__ActorStateTransition__00000003_.Name = `Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy`
	__ActorStateTransition__00000003_.ComputedPrefix = ``
	__ActorStateTransition__00000003_.IsExpanded = false

	__ActorStateTransition__00000004_.Name = `PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd`
	__ActorStateTransition__00000004_.ComputedPrefix = ``
	__ActorStateTransition__00000004_.IsExpanded = false

	__ActorStateTransition__00000005_.Name = `Rick is cynical to Rick is empathic`
	__ActorStateTransition__00000005_.ComputedPrefix = ``
	__ActorStateTransition__00000005_.IsExpanded = false

	__ActorStateTransition__00000006_.Name = `Student 0% to Student 5%`
	__ActorStateTransition__00000006_.ComputedPrefix = ``
	__ActorStateTransition__00000006_.IsExpanded = false

	__ActorStateTransition__00000007_.Name = `Student 5% to Student 15%`
	__ActorStateTransition__00000007_.ComputedPrefix = ``
	__ActorStateTransition__00000007_.IsExpanded = false

	__ActorStateTransitionShape__00000000_.Name = ``
	__ActorStateTransitionShape__00000000_.X = 0.000000
	__ActorStateTransitionShape__00000000_.Y = 0.000000
	__ActorStateTransitionShape__00000000_.Width = 0.000000
	__ActorStateTransitionShape__00000000_.Height = 0.000000
	__ActorStateTransitionShape__00000000_.IsHidden = false

	__ActorStateTransitionShape__00000001_.Name = `A1 to A2`
	__ActorStateTransitionShape__00000001_.X = 0.000000
	__ActorStateTransitionShape__00000001_.Y = 0.000000
	__ActorStateTransitionShape__00000001_.Width = 0.000000
	__ActorStateTransitionShape__00000001_.Height = 0.000000
	__ActorStateTransitionShape__00000001_.IsHidden = false

	__ActorStateTransitionShape__00000002_.Name = `Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy`
	__ActorStateTransitionShape__00000002_.X = 0.000000
	__ActorStateTransitionShape__00000002_.Y = 0.000000
	__ActorStateTransitionShape__00000002_.Width = 0.000000
	__ActorStateTransitionShape__00000002_.Height = 0.000000
	__ActorStateTransitionShape__00000002_.IsHidden = false

	__ActorStateTransitionShape__00000003_.Name = `PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd`
	__ActorStateTransitionShape__00000003_.X = 0.000000
	__ActorStateTransitionShape__00000003_.Y = 0.000000
	__ActorStateTransitionShape__00000003_.Width = 0.000000
	__ActorStateTransitionShape__00000003_.Height = 0.000000
	__ActorStateTransitionShape__00000003_.IsHidden = false

	__ActorStateTransitionShape__00000004_.Name = `Rick is cynical to Rick is empathic`
	__ActorStateTransitionShape__00000004_.X = 0.000000
	__ActorStateTransitionShape__00000004_.Y = 0.000000
	__ActorStateTransitionShape__00000004_.Width = 0.000000
	__ActorStateTransitionShape__00000004_.Height = 0.000000
	__ActorStateTransitionShape__00000004_.IsHidden = false

	__ActorStateTransitionShape__00000005_.Name = `Rick is cynical to Rick is empathic`
	__ActorStateTransitionShape__00000005_.X = 0.000000
	__ActorStateTransitionShape__00000005_.Y = 0.000000
	__ActorStateTransitionShape__00000005_.Width = 0.000000
	__ActorStateTransitionShape__00000005_.Height = 0.000000
	__ActorStateTransitionShape__00000005_.IsHidden = false

	__ActorStateTransitionShape__00000006_.Name = `Student 0% to Student 5%`
	__ActorStateTransitionShape__00000006_.X = 0.000000
	__ActorStateTransitionShape__00000006_.Y = 0.000000
	__ActorStateTransitionShape__00000006_.Width = 0.000000
	__ActorStateTransitionShape__00000006_.Height = 0.000000
	__ActorStateTransitionShape__00000006_.IsHidden = false

	__Analysis__00000000_.Name = `Casablanca`
	__Analysis__00000000_.Description = `A movie in 1942.
With a swedish girl
`
	__Analysis__00000000_.IsScenariosNodeExpanded = false
	__Analysis__00000000_.IsGroupUseNodeExpanded = false
	__Analysis__00000000_.IsGeoObjectUseNodeExpanded = false
	__Analysis__00000000_.IsMapUseNodeExpanded = false
	__Analysis__00000000_.ComputedPrefix = ``
	__Analysis__00000000_.IsExpanded = true

	__Analysis__00000001_.Name = `Horizon(s) analysis`
	__Analysis__00000001_.Description = ``
	__Analysis__00000001_.IsScenariosNodeExpanded = false
	__Analysis__00000001_.IsGroupUseNodeExpanded = false
	__Analysis__00000001_.IsGeoObjectUseNodeExpanded = false
	__Analysis__00000001_.IsMapUseNodeExpanded = false
	__Analysis__00000001_.ComputedPrefix = ``
	__Analysis__00000001_.IsExpanded = false

	__Analysis__00000002_.Name = `Kennan's Long Telegram`
	__Analysis__00000002_.Description = `Content can be found at https://nsarchive2.gwu.edu/coldwar/documents/episode-1/kennan.htm

The "X Article" is an article, formally titled "The Sources of Soviet Conduct", written by George F. Kennan and published under the pseudonym "X" in the July 1947 issue of Foreign Affairs magazine. It widely introduced the term "containment" and advocated for its strategic use against the Soviet Union. It expanded on ideas expressed by Kennan in a confidential February 1946 telegram, formally identified by Kennan's State Department number, "511", but informally dubbed the "long telegram" for its size.

Kennan composed the long telegram to respond to inquiries about the implications of a February 1946 speech by Joseph Stalin.[note 1] Though the speech was in line with previous statements by Stalin, it provoked fear in the American press and public; Time magazine called it "the most warlike pronouncement uttered by any top-rank statesman since V-J Day".[4] The long telegram explained Soviet motivations by recounting the history of Russian rulers as well as the ideology of Marxism–Leninism. It argued that the Soviet leaders used the ideology to characterize the external world as hostile, allowing them to justify their continued hold on power despite a lack of popular support. Washington bureaucrats quickly read the confidential message and accepted it as the best explanation of Soviet behavior. The reception elevated Kennan's reputation within the State Department as one of the government's foremost Soviet experts. 

`
	__Analysis__00000002_.IsScenariosNodeExpanded = false
	__Analysis__00000002_.IsGroupUseNodeExpanded = false
	__Analysis__00000002_.IsGeoObjectUseNodeExpanded = false
	__Analysis__00000002_.IsMapUseNodeExpanded = false
	__Analysis__00000002_.ComputedPrefix = ``
	__Analysis__00000002_.IsExpanded = false

	__Analysis__00000003_.Name = `US Elections`
	__Analysis__00000003_.Description = ``
	__Analysis__00000003_.IsScenariosNodeExpanded = false
	__Analysis__00000003_.IsGroupUseNodeExpanded = false
	__Analysis__00000003_.IsGeoObjectUseNodeExpanded = false
	__Analysis__00000003_.IsMapUseNodeExpanded = false
	__Analysis__00000003_.ComputedPrefix = ``
	__Analysis__00000003_.IsExpanded = false

	__ControlPointShape__00000000_.Name = `Control Point Shape in  0`
	__ControlPointShape__00000000_.X_Relative = 0.116622
	__ControlPointShape__00000000_.Y_Relative = 1.604164
	__ControlPointShape__00000000_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000001_.Name = `Control Point Shape in  0`
	__ControlPointShape__00000001_.X_Relative = -0.255659
	__ControlPointShape__00000001_.Y_Relative = 0.688390
	__ControlPointShape__00000001_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000002_.Name = `Control Point Shape in  1`
	__ControlPointShape__00000002_.X_Relative = 0.086186
	__ControlPointShape__00000002_.Y_Relative = 0.823771
	__ControlPointShape__00000002_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000003_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 0`
	__ControlPointShape__00000003_.X_Relative = 0.645941
	__ControlPointShape__00000003_.Y_Relative = 0.839209
	__ControlPointShape__00000003_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000004_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 1`
	__ControlPointShape__00000004_.X_Relative = 0.024857
	__ControlPointShape__00000004_.Y_Relative = -0.578741
	__ControlPointShape__00000004_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000005_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 2`
	__ControlPointShape__00000005_.X_Relative = -0.099408
	__ControlPointShape__00000005_.Y_Relative = 0.565558
	__ControlPointShape__00000005_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000006_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 3`
	__ControlPointShape__00000006_.X_Relative = 0.039930
	__ControlPointShape__00000006_.Y_Relative = 0.802235
	__ControlPointShape__00000006_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000007_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 3`
	__ControlPointShape__00000007_.X_Relative = -0.095747
	__ControlPointShape__00000007_.Y_Relative = 0.588541
	__ControlPointShape__00000007_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000008_.Name = `Control Point Shape in  2`
	__ControlPointShape__00000008_.X_Relative = 0.397418
	__ControlPointShape__00000008_.Y_Relative = -0.285230
	__ControlPointShape__00000008_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000009_.Name = `Control Point Shape in  2`
	__ControlPointShape__00000009_.X_Relative = 0.089038
	__ControlPointShape__00000009_.Y_Relative = 0.793749
	__ControlPointShape__00000009_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000010_.Name = `Control Point Shape in  2`
	__ControlPointShape__00000010_.X_Relative = 0.089038
	__ControlPointShape__00000010_.Y_Relative = 0.793749
	__ControlPointShape__00000010_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000011_.Name = `Control Point Shape in  0`
	__ControlPointShape__00000011_.X_Relative = 0.611559
	__ControlPointShape__00000011_.Y_Relative = -0.704310
	__ControlPointShape__00000011_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000012_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 1`
	__ControlPointShape__00000012_.X_Relative = 0.082617
	__ControlPointShape__00000012_.Y_Relative = 0.757810
	__ControlPointShape__00000012_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000013_.Name = `Control Point Shape in Political Personality of Soviet Power is dictatorial to PoliticalPersonality ofSoviet Power is amenable to liberal economy 2`
	__ControlPointShape__00000013_.X_Relative = 0.829778
	__ControlPointShape__00000013_.Y_Relative = 1.931247
	__ControlPointShape__00000013_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000014_.Name = `Control Point Shape in  1`
	__ControlPointShape__00000014_.X_Relative = 0.538773
	__ControlPointShape__00000014_.Y_Relative = 1.918747
	__ControlPointShape__00000014_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000015_.Name = `Control Point Shape in  1`
	__ControlPointShape__00000015_.X_Relative = 0.220527
	__ControlPointShape__00000015_.Y_Relative = 0.178898
	__ControlPointShape__00000015_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000016_.Name = `Control Point Shape in  2`
	__ControlPointShape__00000016_.X_Relative = 0.078456
	__ControlPointShape__00000016_.Y_Relative = 0.481248
	__ControlPointShape__00000016_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000017_.Name = `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 0`
	__ControlPointShape__00000017_.X_Relative = 0.924283
	__ControlPointShape__00000017_.Y_Relative = 0.414062
	__ControlPointShape__00000017_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000018_.Name = `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 1`
	__ControlPointShape__00000018_.X_Relative = 1.111109
	__ControlPointShape__00000018_.Y_Relative = -0.594496
	__ControlPointShape__00000018_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000019_.Name = `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 2`
	__ControlPointShape__00000019_.X_Relative = 1.557080
	__ControlPointShape__00000019_.Y_Relative = -0.442795
	__ControlPointShape__00000019_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000020_.Name = `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 3`
	__ControlPointShape__00000020_.X_Relative = 0.140951
	__ControlPointShape__00000020_.Y_Relative = 0.418748
	__ControlPointShape__00000020_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000021_.Name = `Control Point Shape in PoliticalPersonality ofSoviet Power is amenable to liberal economy to Political Personality of Soviet Power is paranoïd 4`
	__ControlPointShape__00000021_.X_Relative = -0.132317
	__ControlPointShape__00000021_.Y_Relative = 0.575401
	__ControlPointShape__00000021_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000022_.Name = `Control Point Shape in Student 0% to Student 5% 0`
	__ControlPointShape__00000022_.X_Relative = 0.577563
	__ControlPointShape__00000022_.Y_Relative = 0.227308
	__ControlPointShape__00000022_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000023_.Name = `Control Point Shape in Student 0% to Student 5% 1`
	__ControlPointShape__00000023_.X_Relative = 0.255789
	__ControlPointShape__00000023_.Y_Relative = 0.381249
	__ControlPointShape__00000023_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000024_.Name = `Control Point Shape in Student 0% to Student 5% 2`
	__ControlPointShape__00000024_.X_Relative = 1.198296
	__ControlPointShape__00000024_.Y_Relative = -1.123532
	__ControlPointShape__00000024_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000025_.Name = `Control Point Shape in Student 0% to Student 5% 3`
	__ControlPointShape__00000025_.X_Relative = -0.703932
	__ControlPointShape__00000025_.Y_Relative = 0.718749
	__ControlPointShape__00000025_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000026_.Name = `Control Point Shape in Student 0% to Student 5% 4`
	__ControlPointShape__00000026_.X_Relative = -1.393083
	__ControlPointShape__00000026_.Y_Relative = 0.356250
	__ControlPointShape__00000026_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000027_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 0`
	__ControlPointShape__00000027_.X_Relative = 0.915168
	__ControlPointShape__00000027_.Y_Relative = 0.556251
	__ControlPointShape__00000027_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000028_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 1`
	__ControlPointShape__00000028_.X_Relative = 1.473502
	__ControlPointShape__00000028_.Y_Relative = -2.081251
	__ControlPointShape__00000028_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000029_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 2`
	__ControlPointShape__00000029_.X_Relative = 0.186697
	__ControlPointShape__00000029_.Y_Relative = 0.274636
	__ControlPointShape__00000029_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000030_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 3`
	__ControlPointShape__00000030_.X_Relative = -0.557143
	__ControlPointShape__00000030_.Y_Relative = 0.395209
	__ControlPointShape__00000030_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000031_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 4`
	__ControlPointShape__00000031_.X_Relative = 1.202668
	__ControlPointShape__00000031_.Y_Relative = -0.068749
	__ControlPointShape__00000031_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000032_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 0`
	__ControlPointShape__00000032_.X_Relative = 0.997263
	__ControlPointShape__00000032_.Y_Relative = 0.595878
	__ControlPointShape__00000032_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000033_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 1`
	__ControlPointShape__00000033_.X_Relative = 1.552819
	__ControlPointShape__00000033_.Y_Relative = -0.694444
	__ControlPointShape__00000033_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000034_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 2`
	__ControlPointShape__00000034_.X_Relative = -0.679298
	__ControlPointShape__00000034_.Y_Relative = 1.034313
	__ControlPointShape__00000034_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000035_.Name = `Control Point Shape in Rick is cynical to Rick is empathic 3`
	__ControlPointShape__00000035_.X_Relative = 0.013615
	__ControlPointShape__00000035_.Y_Relative = 0.401959
	__ControlPointShape__00000035_.IsStartShapeTheClosestShape = false

	__Diagram__00000000_.Name = `Ilsa's trajectory`
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.Description = ``
	__Diagram__00000000_.IsEvolutionDirectionsNodeExpanded = false
	__Diagram__00000000_.IsActorStatesNodeExpanded = false
	__Diagram__00000000_.IsParametersNodeExpanded = false
	__Diagram__00000000_.IsParametersAggregatesNodeExpanded = false
	__Diagram__00000000_.IsActorStateTransitionsNodeExpanded = false
	__Diagram__00000000_.AxisOrign_X = 48.500000
	__Diagram__00000000_.AxisOrign_Y = 230.000000
	__Diagram__00000000_.VerticalAxis_Top_Y = 33.000000
	__Diagram__00000000_.VerticalAxis_Bottom_Y = 523.000000
	__Diagram__00000000_.VerticalAxis_StrokeWidth = 2.000000
	__Diagram__00000000_.HorizontalAxis_Right_X = 689.000000
	__Diagram__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-02-10 01:06:11 +0000 UTC")
	__Diagram__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-02-10 01:06:11 +0000 UTC")
	__Diagram__00000000_.NumberOfYearsBetweenTicks = 0

	__Diagram__00000001_.Name = `New Diagram`
	__Diagram__00000001_.ComputedPrefix = ``
	__Diagram__00000001_.IsExpanded = false
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.Description = ``
	__Diagram__00000001_.IsEvolutionDirectionsNodeExpanded = false
	__Diagram__00000001_.IsActorStatesNodeExpanded = false
	__Diagram__00000001_.IsParametersNodeExpanded = false
	__Diagram__00000001_.IsParametersAggregatesNodeExpanded = true
	__Diagram__00000001_.IsActorStateTransitionsNodeExpanded = true
	__Diagram__00000001_.AxisOrign_X = 25.500000
	__Diagram__00000001_.AxisOrign_Y = 590.000031
	__Diagram__00000001_.VerticalAxis_Top_Y = 16.000000
	__Diagram__00000001_.VerticalAxis_Bottom_Y = 718.000000
	__Diagram__00000001_.VerticalAxis_StrokeWidth = 2.000000
	__Diagram__00000001_.HorizontalAxis_Right_X = 1256.000000
	__Diagram__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1900-01-01 12:59:31 +0000 UTC")
	__Diagram__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1932-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.NumberOfYearsBetweenTicks = 10

	__Diagram__00000002_.Name = `New Diagram`
	__Diagram__00000002_.ComputedPrefix = ``
	__Diagram__00000002_.IsExpanded = false
	__Diagram__00000002_.IsChecked = false
	__Diagram__00000002_.IsShowPrefix = false
	__Diagram__00000002_.Description = ``
	__Diagram__00000002_.IsEvolutionDirectionsNodeExpanded = false
	__Diagram__00000002_.IsActorStatesNodeExpanded = false
	__Diagram__00000002_.IsParametersNodeExpanded = false
	__Diagram__00000002_.IsParametersAggregatesNodeExpanded = false
	__Diagram__00000002_.IsActorStateTransitionsNodeExpanded = false
	__Diagram__00000002_.AxisOrign_X = 53.500000
	__Diagram__00000002_.AxisOrign_Y = 312.000000
	__Diagram__00000002_.VerticalAxis_Top_Y = 33.000000
	__Diagram__00000002_.VerticalAxis_Bottom_Y = 523.000000
	__Diagram__00000002_.VerticalAxis_StrokeWidth = 2.000000
	__Diagram__00000002_.HorizontalAxis_Right_X = 788.000000
	__Diagram__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.NumberOfYearsBetweenTicks = 0

	__Diagram__00000003_.Name = `Reference penetration forcast`
	__Diagram__00000003_.ComputedPrefix = ``
	__Diagram__00000003_.IsExpanded = false
	__Diagram__00000003_.IsChecked = false
	__Diagram__00000003_.IsShowPrefix = false
	__Diagram__00000003_.Description = ``
	__Diagram__00000003_.IsEvolutionDirectionsNodeExpanded = false
	__Diagram__00000003_.IsActorStatesNodeExpanded = false
	__Diagram__00000003_.IsParametersNodeExpanded = false
	__Diagram__00000003_.IsParametersAggregatesNodeExpanded = false
	__Diagram__00000003_.IsActorStateTransitionsNodeExpanded = false
	__Diagram__00000003_.AxisOrign_X = 33.000000
	__Diagram__00000003_.AxisOrign_Y = 427.000000
	__Diagram__00000003_.VerticalAxis_Top_Y = 33.000000
	__Diagram__00000003_.VerticalAxis_Bottom_Y = 535.000000
	__Diagram__00000003_.VerticalAxis_StrokeWidth = 3.000000
	__Diagram__00000003_.HorizontalAxis_Right_X = 1109.000000
	__Diagram__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-12-01 00:00:00 +0000 UTC")
	__Diagram__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-01-01 00:00:00 +0000 UTC")
	__Diagram__00000003_.NumberOfYearsBetweenTicks = 0

	__Diagram__00000004_.Name = `Rick's Trajectory`
	__Diagram__00000004_.ComputedPrefix = ``
	__Diagram__00000004_.IsExpanded = false
	__Diagram__00000004_.IsChecked = true
	__Diagram__00000004_.IsShowPrefix = false
	__Diagram__00000004_.Description = ``
	__Diagram__00000004_.IsEvolutionDirectionsNodeExpanded = false
	__Diagram__00000004_.IsActorStatesNodeExpanded = false
	__Diagram__00000004_.IsParametersNodeExpanded = false
	__Diagram__00000004_.IsParametersAggregatesNodeExpanded = false
	__Diagram__00000004_.IsActorStateTransitionsNodeExpanded = false
	__Diagram__00000004_.AxisOrign_X = 70.000000
	__Diagram__00000004_.AxisOrign_Y = 336.000000
	__Diagram__00000004_.VerticalAxis_Top_Y = 44.000000
	__Diagram__00000004_.VerticalAxis_Bottom_Y = 638.000000
	__Diagram__00000004_.VerticalAxis_StrokeWidth = 3.000000
	__Diagram__00000004_.HorizontalAxis_Right_X = 741.000000
	__Diagram__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1940-12-15 01:06:11 +0000 UTC")
	__Diagram__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1940-12-22 01:06:11 +0000 UTC")
	__Diagram__00000004_.NumberOfYearsBetweenTicks = 0

	__EvolutionDirection__00000000_.Name = `Bottom`
	__EvolutionDirection__00000000_.Description = ``
	__EvolutionDirection__00000000_.ComputedPrefix = ``
	__EvolutionDirection__00000000_.IsExpanded = false

	__EvolutionDirection__00000001_.Name = `Cynical`
	__EvolutionDirection__00000001_.Description = ``
	__EvolutionDirection__00000001_.ComputedPrefix = ``
	__EvolutionDirection__00000001_.IsExpanded = false

	__EvolutionDirection__00000002_.Name = `Democratic`
	__EvolutionDirection__00000002_.Description = ``
	__EvolutionDirection__00000002_.ComputedPrefix = ``
	__EvolutionDirection__00000002_.IsExpanded = false

	__EvolutionDirection__00000003_.Name = `Dictactorship`
	__EvolutionDirection__00000003_.Description = ``
	__EvolutionDirection__00000003_.ComputedPrefix = ``
	__EvolutionDirection__00000003_.IsExpanded = false

	__EvolutionDirection__00000004_.Name = `Empathetic`
	__EvolutionDirection__00000004_.Description = ``
	__EvolutionDirection__00000004_.ComputedPrefix = ``
	__EvolutionDirection__00000004_.IsExpanded = false

	__EvolutionDirection__00000005_.Name = `High Penetration`
	__EvolutionDirection__00000005_.Description = ``
	__EvolutionDirection__00000005_.ComputedPrefix = ``
	__EvolutionDirection__00000005_.IsExpanded = false

	__EvolutionDirection__00000006_.Name = `Low Penetration`
	__EvolutionDirection__00000006_.Description = ``
	__EvolutionDirection__00000006_.ComputedPrefix = ``
	__EvolutionDirection__00000006_.IsExpanded = false

	__EvolutionDirection__00000007_.Name = `Top`
	__EvolutionDirection__00000007_.Description = ``
	__EvolutionDirection__00000007_.ComputedPrefix = ``
	__EvolutionDirection__00000007_.IsExpanded = false

	__EvolutionDirectionShape__00000000_.Name = `Bottom`
	__EvolutionDirectionShape__00000000_.X = 45.000000
	__EvolutionDirectionShape__00000000_.Y = 566.000000
	__EvolutionDirectionShape__00000000_.Width = 112.000000
	__EvolutionDirectionShape__00000000_.Height = 42.000000
	__EvolutionDirectionShape__00000000_.IsHidden = false

	__EvolutionDirectionShape__00000001_.Name = `Cynical`
	__EvolutionDirectionShape__00000001_.X = 73.000000
	__EvolutionDirectionShape__00000001_.Y = 519.999985
	__EvolutionDirectionShape__00000001_.Width = 112.000000
	__EvolutionDirectionShape__00000001_.Height = 42.000000
	__EvolutionDirectionShape__00000001_.IsHidden = false

	__EvolutionDirectionShape__00000002_.Name = `Cynical`
	__EvolutionDirectionShape__00000002_.X = 96.000000
	__EvolutionDirectionShape__00000002_.Y = 658.000000
	__EvolutionDirectionShape__00000002_.Width = 112.000000
	__EvolutionDirectionShape__00000002_.Height = 42.000000
	__EvolutionDirectionShape__00000002_.IsHidden = false

	__EvolutionDirectionShape__00000004_.Name = `Democratic`
	__EvolutionDirectionShape__00000004_.X = 47.000000
	__EvolutionDirectionShape__00000004_.Y = 671.000031
	__EvolutionDirectionShape__00000004_.Width = 112.000000
	__EvolutionDirectionShape__00000004_.Height = 42.000000
	__EvolutionDirectionShape__00000004_.IsHidden = false

	__EvolutionDirectionShape__00000005_.Name = `Dictactorship`
	__EvolutionDirectionShape__00000005_.X = 45.000000
	__EvolutionDirectionShape__00000005_.Y = 16.000008
	__EvolutionDirectionShape__00000005_.Width = 112.000000
	__EvolutionDirectionShape__00000005_.Height = 42.000000
	__EvolutionDirectionShape__00000005_.IsHidden = false

	__EvolutionDirectionShape__00000006_.Name = `Emphatic`
	__EvolutionDirectionShape__00000006_.X = 79.000000
	__EvolutionDirectionShape__00000006_.Y = 50.000000
	__EvolutionDirectionShape__00000006_.Width = 112.000000
	__EvolutionDirectionShape__00000006_.Height = 42.000000
	__EvolutionDirectionShape__00000006_.IsHidden = false

	__EvolutionDirectionShape__00000007_.Name = `Emphatic`
	__EvolutionDirectionShape__00000007_.X = 82.000000
	__EvolutionDirectionShape__00000007_.Y = 21.000015
	__EvolutionDirectionShape__00000007_.Width = 112.000000
	__EvolutionDirectionShape__00000007_.Height = 42.000000
	__EvolutionDirectionShape__00000007_.IsHidden = false

	__EvolutionDirectionShape__00000008_.Name = `High Penetration`
	__EvolutionDirectionShape__00000008_.X = 79.000000
	__EvolutionDirectionShape__00000008_.Y = 50.000000
	__EvolutionDirectionShape__00000008_.Width = 112.000000
	__EvolutionDirectionShape__00000008_.Height = 42.000000
	__EvolutionDirectionShape__00000008_.IsHidden = false

	__EvolutionDirectionShape__00000009_.Name = `Low Penetration`
	__EvolutionDirectionShape__00000009_.X = 69.000000
	__EvolutionDirectionShape__00000009_.Y = 519.000000
	__EvolutionDirectionShape__00000009_.Width = 112.000000
	__EvolutionDirectionShape__00000009_.Height = 42.000000
	__EvolutionDirectionShape__00000009_.IsHidden = false

	__EvolutionDirectionShape__00000010_.Name = `Top`
	__EvolutionDirectionShape__00000010_.X = 83.000000
	__EvolutionDirectionShape__00000010_.Y = 26.000000
	__EvolutionDirectionShape__00000010_.Width = 112.000000
	__EvolutionDirectionShape__00000010_.Height = 42.000000
	__EvolutionDirectionShape__00000010_.IsHidden = false

	__Library__00000000_.Name = ``
	__Library__00000000_.Description = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsAnalysesNodeExpanded = false
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 0.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsExpandedTmp = false

	__Parameter__00000000_.Name = `Abroad capitalistic menace`
	__Parameter__00000000_.Description = ``
	__Parameter__00000000_.IsResponse = false
	__Parameter__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000000_.Force = 0.000000
	__Parameter__00000000_.Tag = ``
	__Parameter__00000000_.ComputedPrefix = ``
	__Parameter__00000000_.IsExpanded = false

	__Parameter__00000001_.Name = `Anti-tsarist rationalization`
	__Parameter__00000001_.Description = ``
	__Parameter__00000001_.IsResponse = false
	__Parameter__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000001_.Force = 0.000000
	__Parameter__00000001_.Tag = ``
	__Parameter__00000001_.ComputedPrefix = ``
	__Parameter__00000001_.IsExpanded = false

	__Parameter__00000002_.Name = `Blogs`
	__Parameter__00000002_.Description = ``
	__Parameter__00000002_.IsResponse = false
	__Parameter__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000002_.Force = 0.000000
	__Parameter__00000002_.Tag = ``
	__Parameter__00000002_.ComputedPrefix = ``
	__Parameter__00000002_.IsExpanded = false

	__Parameter__00000003_.Name = `Capitalism --> Imperialism -> War `
	__Parameter__00000003_.Description = ``
	__Parameter__00000003_.IsResponse = false
	__Parameter__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000003_.Force = 0.000000
	__Parameter__00000003_.Tag = ``
	__Parameter__00000003_.ComputedPrefix = ``
	__Parameter__00000003_.IsExpanded = false

	__Parameter__00000004_.Name = `Communists are a tiny portion of the population`
	__Parameter__00000004_.Description = ``
	__Parameter__00000004_.IsResponse = false
	__Parameter__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000004_.Force = 0.000000
	__Parameter__00000004_.Tag = ``
	__Parameter__00000004_.ComputedPrefix = ``
	__Parameter__00000004_.IsExpanded = false

	__Parameter__00000005_.Name = `Conservatisme`
	__Parameter__00000005_.Description = ``
	__Parameter__00000005_.IsResponse = false
	__Parameter__00000005_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000005_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000005_.Force = 0.000000
	__Parameter__00000005_.Tag = ``
	__Parameter__00000005_.ComputedPrefix = ``
	__Parameter__00000005_.IsExpanded = false

	__Parameter__00000006_.Name = `Ilsa is back in his life ...`
	__Parameter__00000006_.Description = ``
	__Parameter__00000006_.IsResponse = false
	__Parameter__00000006_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000006_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000006_.Force = 0.000000
	__Parameter__00000006_.Tag = ``
	__Parameter__00000006_.ComputedPrefix = ``
	__Parameter__00000006_.IsExpanded = false

	__Parameter__00000007_.Name = `Interventions`
	__Parameter__00000007_.Description = ``
	__Parameter__00000007_.IsResponse = false
	__Parameter__00000007_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000007_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000007_.Force = 0.000000
	__Parameter__00000007_.Tag = ``
	__Parameter__00000007_.ComputedPrefix = ``
	__Parameter__00000007_.IsExpanded = false

	__Parameter__00000008_.Name = `Lenin's Need to protect proletarians from all countries`
	__Parameter__00000008_.Description = ``
	__Parameter__00000008_.IsResponse = false
	__Parameter__00000008_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000008_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000008_.Force = 0.000000
	__Parameter__00000008_.Tag = ``
	__Parameter__00000008_.ComputedPrefix = ``
	__Parameter__00000008_.IsExpanded = false

	__Parameter__00000009_.Name = `Mysteriously gloomy`
	__Parameter__00000009_.Description = ``
	__Parameter__00000009_.IsResponse = false
	__Parameter__00000009_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000009_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000009_.Force = 0.000000
	__Parameter__00000009_.Tag = ``
	__Parameter__00000009_.ComputedPrefix = ``
	__Parameter__00000009_.IsExpanded = false

	__Parameter__00000010_.Name = `P1`
	__Parameter__00000010_.Description = ``
	__Parameter__00000010_.IsResponse = false
	__Parameter__00000010_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000010_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000010_.Force = 0.000000
	__Parameter__00000010_.Tag = ``
	__Parameter__00000010_.ComputedPrefix = ``
	__Parameter__00000010_.IsExpanded = false

	__Parameter__00000011_.Name = `P2`
	__Parameter__00000011_.Description = ``
	__Parameter__00000011_.IsResponse = false
	__Parameter__00000011_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000011_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000011_.Force = 0.000000
	__Parameter__00000011_.Tag = ``
	__Parameter__00000011_.ComputedPrefix = ``
	__Parameter__00000011_.IsExpanded = false

	__Parameter__00000012_.Name = `Paranoïd view of foreign powers (self sustained)`
	__Parameter__00000012_.Description = ``
	__Parameter__00000012_.IsResponse = false
	__Parameter__00000012_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000012_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000012_.Force = 0.000000
	__Parameter__00000012_.Tag = ``
	__Parameter__00000012_.ComputedPrefix = ``
	__Parameter__00000012_.IsExpanded = false

	__Parameter__00000013_.Name = `Publications`
	__Parameter__00000013_.Description = ``
	__Parameter__00000013_.IsResponse = false
	__Parameter__00000013_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000013_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000013_.Force = 0.000000
	__Parameter__00000013_.Tag = ``
	__Parameter__00000013_.ComputedPrefix = ``
	__Parameter__00000013_.IsExpanded = false

	__Parameter__00000014_.Name = `Rejection of early war communism`
	__Parameter__00000014_.Description = ``
	__Parameter__00000014_.IsResponse = false
	__Parameter__00000014_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000014_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000014_.Force = 0.000000
	__Parameter__00000014_.Tag = ``
	__Parameter__00000014_.ComputedPrefix = ``
	__Parameter__00000014_.IsExpanded = false

	__Parameter__00000015_.Name = `Russian Asiatic attraction to imperialism`
	__Parameter__00000015_.Description = ``
	__Parameter__00000015_.IsResponse = false
	__Parameter__00000015_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000015_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000015_.Force = 0.000000
	__Parameter__00000015_.Tag = ``
	__Parameter__00000015_.ComputedPrefix = ``
	__Parameter__00000015_.IsExpanded = false

	__Parameter__00000016_.Name = `Siloviki's incitation to maintain internal oppression`
	__Parameter__00000016_.Description = ``
	__Parameter__00000016_.IsResponse = false
	__Parameter__00000016_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000016_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000016_.Force = 0.000000
	__Parameter__00000016_.Tag = ``
	__Parameter__00000016_.ComputedPrefix = ``
	__Parameter__00000016_.IsExpanded = false

	__Parameter__00000017_.Name = `Stalin's intolerance to power sharing`
	__Parameter__00000017_.Description = ``
	__Parameter__00000017_.IsResponse = false
	__Parameter__00000017_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000017_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Parameter__00000017_.Force = 0.000000
	__Parameter__00000017_.Tag = ``
	__Parameter__00000017_.ComputedPrefix = ``
	__Parameter__00000017_.IsExpanded = false

	__ParameterShape__00000000_.Name = `Blogs`
	__ParameterShape__00000000_.Direction = models.DIRECTION_UP
	__ParameterShape__00000000_.ShapeIsComputedFromModel = false
	__ParameterShape__00000000_.X = 496.000000
	__ParameterShape__00000000_.Y = 186.000015
	__ParameterShape__00000000_.Width = 139.500000
	__ParameterShape__00000000_.Height = 72.000000
	__ParameterShape__00000000_.IsHidden = false

	__ParameterShape__00000001_.Name = `Capitalism --> Imperialism -> War `
	__ParameterShape__00000001_.Direction = models.DIRECTION_UP
	__ParameterShape__00000001_.ShapeIsComputedFromModel = false
	__ParameterShape__00000001_.X = 234.000000
	__ParameterShape__00000001_.Y = 386.999985
	__ParameterShape__00000001_.Width = 180.000000
	__ParameterShape__00000001_.Height = 80.000000
	__ParameterShape__00000001_.IsHidden = false

	__ParameterShape__00000002_.Name = `Circumstances`
	__ParameterShape__00000002_.Direction = models.DIRECTION_UP
	__ParameterShape__00000002_.ShapeIsComputedFromModel = false
	__ParameterShape__00000002_.X = 94.000000
	__ParameterShape__00000002_.Y = 470.000016
	__ParameterShape__00000002_.Width = 133.000000
	__ParameterShape__00000002_.Height = 107.111111
	__ParameterShape__00000002_.IsHidden = false

	__ParameterShape__00000003_.Name = `Communists are a tiny portion of the population`
	__ParameterShape__00000003_.Direction = models.DIRECTION_UP
	__ParameterShape__00000003_.ShapeIsComputedFromModel = false
	__ParameterShape__00000003_.X = 173.000000
	__ParameterShape__00000003_.Y = 26.000000
	__ParameterShape__00000003_.Width = 180.000000
	__ParameterShape__00000003_.Height = 80.000000
	__ParameterShape__00000003_.IsHidden = false

	__ParameterShape__00000004_.Name = `Conservatisme`
	__ParameterShape__00000004_.Direction = models.DIRECTION_DOWN
	__ParameterShape__00000004_.ShapeIsComputedFromModel = false
	__ParameterShape__00000004_.X = 249.000000
	__ParameterShape__00000004_.Y = 61.000000
	__ParameterShape__00000004_.Width = 147.750000
	__ParameterShape__00000004_.Height = 79.000000
	__ParameterShape__00000004_.IsHidden = false

	__ParameterShape__00000006_.Name = `Foreign capitalistic menace`
	__ParameterShape__00000006_.Direction = models.DIRECTION_UP
	__ParameterShape__00000006_.ShapeIsComputedFromModel = false
	__ParameterShape__00000006_.X = 808.000000
	__ParameterShape__00000006_.Y = 487.000031
	__ParameterShape__00000006_.Width = 180.000000
	__ParameterShape__00000006_.Height = 80.000000
	__ParameterShape__00000006_.IsHidden = false

	__ParameterShape__00000007_.Name = `Ilsa is back in his life`
	__ParameterShape__00000007_.Direction = models.DIRECTION_UP
	__ParameterShape__00000007_.ShapeIsComputedFromModel = false
	__ParameterShape__00000007_.X = 478.000000
	__ParameterShape__00000007_.Y = 282.000000
	__ParameterShape__00000007_.Width = 180.000000
	__ParameterShape__00000007_.Height = 80.000000
	__ParameterShape__00000007_.IsHidden = false

	__ParameterShape__00000008_.Name = `Ilsa is bak in his life`
	__ParameterShape__00000008_.Direction = models.DIRECTION_UP
	__ParameterShape__00000008_.ShapeIsComputedFromModel = false
	__ParameterShape__00000008_.X = 443.000000
	__ParameterShape__00000008_.Y = 420.000000
	__ParameterShape__00000008_.Width = 244.000000
	__ParameterShape__00000008_.Height = 75.724138
	__ParameterShape__00000008_.IsHidden = false

	__ParameterShape__00000009_.Name = `Interventions`
	__ParameterShape__00000009_.Direction = models.DIRECTION_UP
	__ParameterShape__00000009_.ShapeIsComputedFromModel = false
	__ParameterShape__00000009_.X = 758.000000
	__ParameterShape__00000009_.Y = 199.000015
	__ParameterShape__00000009_.Width = 137.000000
	__ParameterShape__00000009_.Height = 67.888889
	__ParameterShape__00000009_.IsHidden = false

	__ParameterShape__00000010_.Name = `Lenin's Need to protect proletarians from all countries`
	__ParameterShape__00000010_.Direction = models.DIRECTION_UP
	__ParameterShape__00000010_.ShapeIsComputedFromModel = false
	__ParameterShape__00000010_.X = 433.000000
	__ParameterShape__00000010_.Y = 465.000000
	__ParameterShape__00000010_.Width = 180.000000
	__ParameterShape__00000010_.Height = 80.000000
	__ParameterShape__00000010_.IsHidden = false

	__ParameterShape__00000012_.Name = `Mysteriously gloomy`
	__ParameterShape__00000012_.Direction = models.DIRECTION_DOWN
	__ParameterShape__00000012_.ShapeIsComputedFromModel = false
	__ParameterShape__00000012_.X = 131.000000
	__ParameterShape__00000012_.Y = 266.000000
	__ParameterShape__00000012_.Width = 180.000000
	__ParameterShape__00000012_.Height = 80.000000
	__ParameterShape__00000012_.IsHidden = false

	__ParameterShape__00000013_.Name = `P1`
	__ParameterShape__00000013_.Direction = models.DIRECTION_DOWN
	__ParameterShape__00000013_.ShapeIsComputedFromModel = false
	__ParameterShape__00000013_.X = 230.000000
	__ParameterShape__00000013_.Y = 104.000000
	__ParameterShape__00000013_.Width = 180.000000
	__ParameterShape__00000013_.Height = 80.000000
	__ParameterShape__00000013_.IsHidden = false

	__ParameterShape__00000014_.Name = `P2`
	__ParameterShape__00000014_.Direction = models.DIRECTION_UP
	__ParameterShape__00000014_.ShapeIsComputedFromModel = false
	__ParameterShape__00000014_.X = 640.000000
	__ParameterShape__00000014_.Y = 340.000000
	__ParameterShape__00000014_.Width = 180.000000
	__ParameterShape__00000014_.Height = 80.000000
	__ParameterShape__00000014_.IsHidden = false

	__ParameterShape__00000015_.Name = `Paranoïd view of foreign powers (self sustained)`
	__ParameterShape__00000015_.Direction = models.DIRECTION_UP
	__ParameterShape__00000015_.ShapeIsComputedFromModel = false
	__ParameterShape__00000015_.X = 896.000000
	__ParameterShape__00000015_.Y = 163.000015
	__ParameterShape__00000015_.Width = 180.000000
	__ParameterShape__00000015_.Height = 80.000000
	__ParameterShape__00000015_.IsHidden = false

	__ParameterShape__00000016_.Name = `Publications`
	__ParameterShape__00000016_.Direction = models.DIRECTION_UP
	__ParameterShape__00000016_.ShapeIsComputedFromModel = false
	__ParameterShape__00000016_.X = 325.000000
	__ParameterShape__00000016_.Y = 289.000000
	__ParameterShape__00000016_.Width = 137.000000
	__ParameterShape__00000016_.Height = 60.888889
	__ParameterShape__00000016_.IsHidden = false

	__ParameterShape__00000017_.Name = `Rejection of early war communism`
	__ParameterShape__00000017_.Direction = models.DIRECTION_DOWN
	__ParameterShape__00000017_.ShapeIsComputedFromModel = false
	__ParameterShape__00000017_.X = 455.000000
	__ParameterShape__00000017_.Y = 118.999992
	__ParameterShape__00000017_.Width = 160.000000
	__ParameterShape__00000017_.Height = 67.000000
	__ParameterShape__00000017_.IsHidden = false

	__ParameterShape__00000018_.Name = `Russian Asiatic attraction to imperialism`
	__ParameterShape__00000018_.Direction = models.DIRECTION_UP
	__ParameterShape__00000018_.ShapeIsComputedFromModel = false
	__ParameterShape__00000018_.X = 747.000000
	__ParameterShape__00000018_.Y = 248.000000
	__ParameterShape__00000018_.Width = 180.000000
	__ParameterShape__00000018_.Height = 80.000000
	__ParameterShape__00000018_.IsHidden = false

	__ParameterShape__00000019_.Name = `Siloviki's incitation to maintain internal oppression`
	__ParameterShape__00000019_.Direction = models.DIRECTION_UP
	__ParameterShape__00000019_.ShapeIsComputedFromModel = false
	__ParameterShape__00000019_.X = 1079.000000
	__ParameterShape__00000019_.Y = 416.000000
	__ParameterShape__00000019_.Width = 180.000000
	__ParameterShape__00000019_.Height = 80.000000
	__ParameterShape__00000019_.IsHidden = false

	__ParameterShape__00000020_.Name = `Stalin's intolerance to power sharing`
	__ParameterShape__00000020_.Direction = models.DIRECTION_UP
	__ParameterShape__00000020_.ShapeIsComputedFromModel = false
	__ParameterShape__00000020_.X = 649.000000
	__ParameterShape__00000020_.Y = 371.000000
	__ParameterShape__00000020_.Width = 180.000000
	__ParameterShape__00000020_.Height = 80.000000
	__ParameterShape__00000020_.IsHidden = false

	__ParametersAggregate__00000000_.Name = ``
	__ParametersAggregate__00000000_.Tag = ``
	__ParametersAggregate__00000000_.Description = ``
	__ParametersAggregate__00000000_.ComputedPrefix = ``
	__ParametersAggregate__00000000_.IsExpanded = false

	__ParametersAggregate__00000001_.Name = `Circumstances`
	__ParametersAggregate__00000001_.Tag = ``
	__ParametersAggregate__00000001_.Description = ``
	__ParametersAggregate__00000001_.ComputedPrefix = ``
	__ParametersAggregate__00000001_.IsExpanded = false

	__ParametersAggregate__00000002_.Name = `Evangelisation`
	__ParametersAggregate__00000002_.Tag = ``
	__ParametersAggregate__00000002_.Description = ``
	__ParametersAggregate__00000002_.ComputedPrefix = ``
	__ParametersAggregate__00000002_.IsExpanded = false

	__ParametersAggregate__00000003_.Name = `Marxian ideology`
	__ParametersAggregate__00000003_.Tag = ``
	__ParametersAggregate__00000003_.Description = ``
	__ParametersAggregate__00000003_.ComputedPrefix = ``
	__ParametersAggregate__00000003_.IsExpanded = false

	__ParametersAggregate__00000004_.Name = `P1+P2`
	__ParametersAggregate__00000004_.Tag = ``
	__ParametersAggregate__00000004_.Description = ``
	__ParametersAggregate__00000004_.ComputedPrefix = ``
	__ParametersAggregate__00000004_.IsExpanded = false

	__ParametersAggregateShape__00000000_.Name = ``
	__ParametersAggregateShape__00000000_.Direction = ""
	__ParametersAggregateShape__00000000_.X = 695.000000
	__ParametersAggregateShape__00000000_.Y = 26.000000
	__ParametersAggregateShape__00000000_.Width = 180.000000
	__ParametersAggregateShape__00000000_.Height = 80.000000
	__ParametersAggregateShape__00000000_.IsHidden = false

	__ParametersAggregateShape__00000001_.Name = `Circumstances`
	__ParametersAggregateShape__00000001_.Direction = ""
	__ParametersAggregateShape__00000001_.X = 677.000000
	__ParametersAggregateShape__00000001_.Y = 9.999962
	__ParametersAggregateShape__00000001_.Width = 193.500000
	__ParametersAggregateShape__00000001_.Height = 58.000000
	__ParametersAggregateShape__00000001_.IsHidden = false

	__ParametersAggregateShape__00000003_.Name = `Evangelisation`
	__ParametersAggregateShape__00000003_.Direction = models.DIRECTION_UP
	__ParametersAggregateShape__00000003_.X = 524.000000
	__ParametersAggregateShape__00000003_.Y = 337.000000
	__ParametersAggregateShape__00000003_.Width = 180.000000
	__ParametersAggregateShape__00000003_.Height = 73.000000
	__ParametersAggregateShape__00000003_.IsHidden = false

	__ParametersAggregateShape__00000004_.Name = `Marxian ideology`
	__ParametersAggregateShape__00000004_.Direction = models.DIRECTION_UP
	__ParametersAggregateShape__00000004_.X = 432.000000
	__ParametersAggregateShape__00000004_.Y = 621.999985
	__ParametersAggregateShape__00000004_.Width = 123.750000
	__ParametersAggregateShape__00000004_.Height = 66.000000
	__ParametersAggregateShape__00000004_.IsHidden = false

	__Scenario__00000000_.Name = `Kennan's analysis`
	__Scenario__00000000_.Description = ``
	__Scenario__00000000_.IsDiagramsNodeExpanded = true
	__Scenario__00000000_.IsActorStatesNodeExpanded = false
	__Scenario__00000000_.IsActorStateTransitionsNodeExpanded = false
	__Scenario__00000000_.IsEvolutionDirectionsNodeExpanded = false
	__Scenario__00000000_.IsParametersNodeExpanded = false
	__Scenario__00000000_.IsParametersAggretatesNodeExpanded = false
	__Scenario__00000000_.ComputedPrefix = ``
	__Scenario__00000000_.IsExpanded = false

	__Scenario__00000001_.Name = `Nicky Halley trajectory`
	__Scenario__00000001_.Description = ``
	__Scenario__00000001_.IsDiagramsNodeExpanded = true
	__Scenario__00000001_.IsActorStatesNodeExpanded = false
	__Scenario__00000001_.IsActorStateTransitionsNodeExpanded = false
	__Scenario__00000001_.IsEvolutionDirectionsNodeExpanded = false
	__Scenario__00000001_.IsParametersNodeExpanded = false
	__Scenario__00000001_.IsParametersAggretatesNodeExpanded = false
	__Scenario__00000001_.ComputedPrefix = ``
	__Scenario__00000001_.IsExpanded = false

	__Scenario__00000002_.Name = `Penetration Forcast`
	__Scenario__00000002_.Description = ``
	__Scenario__00000002_.IsDiagramsNodeExpanded = true
	__Scenario__00000002_.IsActorStatesNodeExpanded = false
	__Scenario__00000002_.IsActorStateTransitionsNodeExpanded = false
	__Scenario__00000002_.IsEvolutionDirectionsNodeExpanded = false
	__Scenario__00000002_.IsParametersNodeExpanded = false
	__Scenario__00000002_.IsParametersAggretatesNodeExpanded = false
	__Scenario__00000002_.ComputedPrefix = ``
	__Scenario__00000002_.IsExpanded = false

	__Scenario__00000003_.Name = `Scenario avec intervention sociale concertée`
	__Scenario__00000003_.Description = ``
	__Scenario__00000003_.IsDiagramsNodeExpanded = false
	__Scenario__00000003_.IsActorStatesNodeExpanded = false
	__Scenario__00000003_.IsActorStateTransitionsNodeExpanded = false
	__Scenario__00000003_.IsEvolutionDirectionsNodeExpanded = false
	__Scenario__00000003_.IsParametersNodeExpanded = false
	__Scenario__00000003_.IsParametersAggretatesNodeExpanded = false
	__Scenario__00000003_.ComputedPrefix = ``
	__Scenario__00000003_.IsExpanded = false

	__Scenario__00000004_.Name = `Scénario d'évolution multi domaines sans intervention`
	__Scenario__00000004_.Description = ``
	__Scenario__00000004_.IsDiagramsNodeExpanded = false
	__Scenario__00000004_.IsActorStatesNodeExpanded = false
	__Scenario__00000004_.IsActorStateTransitionsNodeExpanded = false
	__Scenario__00000004_.IsEvolutionDirectionsNodeExpanded = false
	__Scenario__00000004_.IsParametersNodeExpanded = false
	__Scenario__00000004_.IsParametersAggretatesNodeExpanded = false
	__Scenario__00000004_.ComputedPrefix = ``
	__Scenario__00000004_.IsExpanded = false

	__Scenario__00000005_.Name = `The Movie`
	__Scenario__00000005_.Description = ``
	__Scenario__00000005_.IsDiagramsNodeExpanded = true
	__Scenario__00000005_.IsActorStatesNodeExpanded = false
	__Scenario__00000005_.IsActorStateTransitionsNodeExpanded = false
	__Scenario__00000005_.IsEvolutionDirectionsNodeExpanded = false
	__Scenario__00000005_.IsParametersNodeExpanded = false
	__Scenario__00000005_.IsParametersAggretatesNodeExpanded = false
	__Scenario__00000005_.ComputedPrefix = ``
	__Scenario__00000005_.IsExpanded = true

	__Workspace__00000000_.Name = `Default`
	__Workspace__00000000_.ComputedPrefix = ``
	__Workspace__00000000_.IsExpanded = false

	// insertion point for setup of pointers
	__ActorStateShape__00000004_.ActorState = __ActorState__00000000_
	__ActorStateShape__00000005_.ActorState = __ActorState__00000001_
	__ActorStateShape__00000007_.ActorState = __ActorState__00000003_
	__ActorStateShape__00000008_.ActorState = __ActorState__00000004_
	__ActorStateShape__00000009_.ActorState = __ActorState__00000005_
	__ActorStateShape__00000010_.ActorState = __ActorState__00000002_
	__ActorStateShape__00000011_.ActorState = __ActorState__00000006_
	__ActorStateShape__00000012_.ActorState = __ActorState__00000006_
	__ActorStateShape__00000013_.ActorState = __ActorState__00000007_
	__ActorStateShape__00000014_.ActorState = __ActorState__00000007_
	__ActorStateShape__00000015_.ActorState = __ActorState__00000008_
	__ActorStateShape__00000016_.ActorState = __ActorState__00000009_
	__ActorStateTransition__00000000_.StartState = __ActorState__00000005_
	__ActorStateTransition__00000000_.EndState = __ActorState__00000003_
	__ActorStateTransition__00000001_.StartState = __ActorState__00000000_
	__ActorStateTransition__00000001_.EndState = __ActorState__00000001_
	__ActorStateTransition__00000002_.StartState = nil
	__ActorStateTransition__00000002_.EndState = __ActorState__00000001_
	__ActorStateTransition__00000003_.StartState = __ActorState__00000003_
	__ActorStateTransition__00000003_.EndState = __ActorState__00000002_
	__ActorStateTransition__00000004_.StartState = __ActorState__00000002_
	__ActorStateTransition__00000004_.EndState = __ActorState__00000004_
	__ActorStateTransition__00000005_.StartState = __ActorState__00000006_
	__ActorStateTransition__00000005_.EndState = __ActorState__00000007_
	__ActorStateTransition__00000006_.StartState = __ActorState__00000008_
	__ActorStateTransition__00000006_.EndState = __ActorState__00000009_
	__ActorStateTransition__00000007_.StartState = __ActorState__00000009_
	__ActorStateTransition__00000007_.EndState = nil
	__ActorStateTransitionShape__00000000_.ActorStateTransition = __ActorStateTransition__00000000_
	__ActorStateTransitionShape__00000000_.Start = __ActorStateShape__00000009_
	__ActorStateTransitionShape__00000000_.End = __ActorStateShape__00000007_
	__ActorStateTransitionShape__00000000_.ControlPointShapes = append(__ActorStateTransitionShape__00000000_.ControlPointShapes, __ControlPointShape__00000015_)
	__ActorStateTransitionShape__00000000_.ControlPointShapes = append(__ActorStateTransitionShape__00000000_.ControlPointShapes, __ControlPointShape__00000011_)
	__ActorStateTransitionShape__00000000_.ControlPointShapes = append(__ActorStateTransitionShape__00000000_.ControlPointShapes, __ControlPointShape__00000016_)
	__ActorStateTransitionShape__00000001_.ActorStateTransition = __ActorStateTransition__00000001_
	__ActorStateTransitionShape__00000001_.Start = __ActorStateShape__00000004_
	__ActorStateTransitionShape__00000001_.End = __ActorStateShape__00000005_
	__ActorStateTransitionShape__00000002_.ActorStateTransition = __ActorStateTransition__00000003_
	__ActorStateTransitionShape__00000002_.Start = __ActorStateShape__00000007_
	__ActorStateTransitionShape__00000002_.End = __ActorStateShape__00000010_
	__ActorStateTransitionShape__00000002_.ControlPointShapes = append(__ActorStateTransitionShape__00000002_.ControlPointShapes, __ControlPointShape__00000003_)
	__ActorStateTransitionShape__00000002_.ControlPointShapes = append(__ActorStateTransitionShape__00000002_.ControlPointShapes, __ControlPointShape__00000013_)
	__ActorStateTransitionShape__00000002_.ControlPointShapes = append(__ActorStateTransitionShape__00000002_.ControlPointShapes, __ControlPointShape__00000012_)
	__ActorStateTransitionShape__00000003_.ActorStateTransition = __ActorStateTransition__00000004_
	__ActorStateTransitionShape__00000003_.Start = __ActorStateShape__00000010_
	__ActorStateTransitionShape__00000003_.End = __ActorStateShape__00000008_
	__ActorStateTransitionShape__00000003_.ControlPointShapes = append(__ActorStateTransitionShape__00000003_.ControlPointShapes, __ControlPointShape__00000017_)
	__ActorStateTransitionShape__00000003_.ControlPointShapes = append(__ActorStateTransitionShape__00000003_.ControlPointShapes, __ControlPointShape__00000018_)
	__ActorStateTransitionShape__00000003_.ControlPointShapes = append(__ActorStateTransitionShape__00000003_.ControlPointShapes, __ControlPointShape__00000019_)
	__ActorStateTransitionShape__00000003_.ControlPointShapes = append(__ActorStateTransitionShape__00000003_.ControlPointShapes, __ControlPointShape__00000021_)
	__ActorStateTransitionShape__00000003_.ControlPointShapes = append(__ActorStateTransitionShape__00000003_.ControlPointShapes, __ControlPointShape__00000020_)
	__ActorStateTransitionShape__00000004_.ActorStateTransition = __ActorStateTransition__00000005_
	__ActorStateTransitionShape__00000004_.Start = __ActorStateShape__00000012_
	__ActorStateTransitionShape__00000004_.End = __ActorStateShape__00000013_
	__ActorStateTransitionShape__00000004_.ControlPointShapes = append(__ActorStateTransitionShape__00000004_.ControlPointShapes, __ControlPointShape__00000027_)
	__ActorStateTransitionShape__00000004_.ControlPointShapes = append(__ActorStateTransitionShape__00000004_.ControlPointShapes, __ControlPointShape__00000031_)
	__ActorStateTransitionShape__00000004_.ControlPointShapes = append(__ActorStateTransitionShape__00000004_.ControlPointShapes, __ControlPointShape__00000028_)
	__ActorStateTransitionShape__00000004_.ControlPointShapes = append(__ActorStateTransitionShape__00000004_.ControlPointShapes, __ControlPointShape__00000030_)
	__ActorStateTransitionShape__00000004_.ControlPointShapes = append(__ActorStateTransitionShape__00000004_.ControlPointShapes, __ControlPointShape__00000029_)
	__ActorStateTransitionShape__00000005_.ActorStateTransition = __ActorStateTransition__00000005_
	__ActorStateTransitionShape__00000005_.Start = __ActorStateShape__00000011_
	__ActorStateTransitionShape__00000005_.End = __ActorStateShape__00000014_
	__ActorStateTransitionShape__00000005_.ControlPointShapes = append(__ActorStateTransitionShape__00000005_.ControlPointShapes, __ControlPointShape__00000032_)
	__ActorStateTransitionShape__00000005_.ControlPointShapes = append(__ActorStateTransitionShape__00000005_.ControlPointShapes, __ControlPointShape__00000033_)
	__ActorStateTransitionShape__00000005_.ControlPointShapes = append(__ActorStateTransitionShape__00000005_.ControlPointShapes, __ControlPointShape__00000034_)
	__ActorStateTransitionShape__00000005_.ControlPointShapes = append(__ActorStateTransitionShape__00000005_.ControlPointShapes, __ControlPointShape__00000035_)
	__ActorStateTransitionShape__00000006_.ActorStateTransition = __ActorStateTransition__00000006_
	__ActorStateTransitionShape__00000006_.Start = __ActorStateShape__00000015_
	__ActorStateTransitionShape__00000006_.End = __ActorStateShape__00000016_
	__ActorStateTransitionShape__00000006_.ControlPointShapes = append(__ActorStateTransitionShape__00000006_.ControlPointShapes, __ControlPointShape__00000022_)
	__ActorStateTransitionShape__00000006_.ControlPointShapes = append(__ActorStateTransitionShape__00000006_.ControlPointShapes, __ControlPointShape__00000024_)
	__ActorStateTransitionShape__00000006_.ControlPointShapes = append(__ActorStateTransitionShape__00000006_.ControlPointShapes, __ControlPointShape__00000026_)
	__ActorStateTransitionShape__00000006_.ControlPointShapes = append(__ActorStateTransitionShape__00000006_.ControlPointShapes, __ControlPointShape__00000023_)
	__Analysis__00000000_.Scenarios = append(__Analysis__00000000_.Scenarios, __Scenario__00000005_)
	__Analysis__00000001_.Scenarios = append(__Analysis__00000001_.Scenarios, __Scenario__00000002_)
	__Analysis__00000002_.Scenarios = append(__Analysis__00000002_.Scenarios, __Scenario__00000000_)
	__Analysis__00000003_.Scenarios = append(__Analysis__00000003_.Scenarios, __Scenario__00000001_)
	__Diagram__00000000_.EvolutionDirectionShapes = append(__Diagram__00000000_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000001_)
	__Diagram__00000000_.EvolutionDirectionShapes = append(__Diagram__00000000_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000007_)
	__Diagram__00000000_.ActorStateShapes = append(__Diagram__00000000_.ActorStateShapes, __ActorStateShape__00000013_)
	__Diagram__00000000_.ActorStateShapes = append(__Diagram__00000000_.ActorStateShapes, __ActorStateShape__00000012_)
	__Diagram__00000000_.ParameterShapes = append(__Diagram__00000000_.ParameterShapes, __ParameterShape__00000007_)
	__Diagram__00000000_.ParameterShapes = append(__Diagram__00000000_.ParameterShapes, __ParameterShape__00000012_)
	__Diagram__00000000_.ActorStateTransitionShapes = append(__Diagram__00000000_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000004_)
	__Diagram__00000001_.EvolutionDirectionShapes = append(__Diagram__00000001_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000005_)
	__Diagram__00000001_.EvolutionDirectionShapes = append(__Diagram__00000001_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000004_)
	__Diagram__00000001_.ActorStateShapes = append(__Diagram__00000001_.ActorStateShapes, __ActorStateShape__00000009_)
	__Diagram__00000001_.ActorStateShapes = append(__Diagram__00000001_.ActorStateShapes, __ActorStateShape__00000007_)
	__Diagram__00000001_.ActorStateShapes = append(__Diagram__00000001_.ActorStateShapes, __ActorStateShape__00000010_)
	__Diagram__00000001_.ActorStateShapes = append(__Diagram__00000001_.ActorStateShapes, __ActorStateShape__00000008_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000002_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000001_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000003_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000017_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000020_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000018_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000015_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000006_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000010_)
	__Diagram__00000001_.ParameterShapes = append(__Diagram__00000001_.ParameterShapes, __ParameterShape__00000019_)
	__Diagram__00000001_.ScenarioParameterShapes = append(__Diagram__00000001_.ScenarioParameterShapes, __ParametersAggregateShape__00000004_)
	__Diagram__00000001_.ScenarioParameterShapes = append(__Diagram__00000001_.ScenarioParameterShapes, __ParametersAggregateShape__00000001_)
	__Diagram__00000001_.ActorStateTransitionShapes = append(__Diagram__00000001_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000000_)
	__Diagram__00000001_.ActorStateTransitionShapes = append(__Diagram__00000001_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000003_)
	__Diagram__00000001_.ActorStateTransitionShapes = append(__Diagram__00000001_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000002_)
	__Diagram__00000002_.EvolutionDirectionShapes = append(__Diagram__00000002_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000010_)
	__Diagram__00000002_.EvolutionDirectionShapes = append(__Diagram__00000002_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000000_)
	__Diagram__00000002_.ActorStateShapes = append(__Diagram__00000002_.ActorStateShapes, __ActorStateShape__00000005_)
	__Diagram__00000002_.ActorStateShapes = append(__Diagram__00000002_.ActorStateShapes, __ActorStateShape__00000004_)
	__Diagram__00000002_.ParameterShapes = append(__Diagram__00000002_.ParameterShapes, __ParameterShape__00000013_)
	__Diagram__00000002_.ParameterShapes = append(__Diagram__00000002_.ParameterShapes, __ParameterShape__00000014_)
	__Diagram__00000002_.ScenarioParameterShapes = append(__Diagram__00000002_.ScenarioParameterShapes, __ParametersAggregateShape__00000000_)
	__Diagram__00000002_.ActorStateTransitionShapes = append(__Diagram__00000002_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000001_)
	__Diagram__00000003_.EvolutionDirectionShapes = append(__Diagram__00000003_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000008_)
	__Diagram__00000003_.EvolutionDirectionShapes = append(__Diagram__00000003_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000009_)
	__Diagram__00000003_.ActorStateShapes = append(__Diagram__00000003_.ActorStateShapes, __ActorStateShape__00000015_)
	__Diagram__00000003_.ActorStateShapes = append(__Diagram__00000003_.ActorStateShapes, __ActorStateShape__00000016_)
	__Diagram__00000003_.ParameterShapes = append(__Diagram__00000003_.ParameterShapes, __ParameterShape__00000016_)
	__Diagram__00000003_.ParameterShapes = append(__Diagram__00000003_.ParameterShapes, __ParameterShape__00000009_)
	__Diagram__00000003_.ParameterShapes = append(__Diagram__00000003_.ParameterShapes, __ParameterShape__00000000_)
	__Diagram__00000003_.ParameterShapes = append(__Diagram__00000003_.ParameterShapes, __ParameterShape__00000004_)
	__Diagram__00000003_.ScenarioParameterShapes = append(__Diagram__00000003_.ScenarioParameterShapes, __ParametersAggregateShape__00000003_)
	__Diagram__00000003_.ActorStateTransitionShapes = append(__Diagram__00000003_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000006_)
	__Diagram__00000004_.EvolutionDirectionShapes = append(__Diagram__00000004_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000006_)
	__Diagram__00000004_.EvolutionDirectionShapes = append(__Diagram__00000004_.EvolutionDirectionShapes, __EvolutionDirectionShape__00000002_)
	__Diagram__00000004_.ActorStateShapes = append(__Diagram__00000004_.ActorStateShapes, __ActorStateShape__00000014_)
	__Diagram__00000004_.ActorStateShapes = append(__Diagram__00000004_.ActorStateShapes, __ActorStateShape__00000011_)
	__Diagram__00000004_.ParameterShapes = append(__Diagram__00000004_.ParameterShapes, __ParameterShape__00000008_)
	__Diagram__00000004_.ActorStateTransitionShapes = append(__Diagram__00000004_.ActorStateTransitionShapes, __ActorStateTransitionShape__00000005_)
	__EvolutionDirectionShape__00000000_.EvolutionDirection = __EvolutionDirection__00000000_
	__EvolutionDirectionShape__00000001_.EvolutionDirection = __EvolutionDirection__00000001_
	__EvolutionDirectionShape__00000002_.EvolutionDirection = __EvolutionDirection__00000001_
	__EvolutionDirectionShape__00000004_.EvolutionDirection = __EvolutionDirection__00000002_
	__EvolutionDirectionShape__00000005_.EvolutionDirection = __EvolutionDirection__00000003_
	__EvolutionDirectionShape__00000006_.EvolutionDirection = __EvolutionDirection__00000004_
	__EvolutionDirectionShape__00000007_.EvolutionDirection = __EvolutionDirection__00000004_
	__EvolutionDirectionShape__00000008_.EvolutionDirection = __EvolutionDirection__00000005_
	__EvolutionDirectionShape__00000009_.EvolutionDirection = __EvolutionDirection__00000006_
	__EvolutionDirectionShape__00000010_.EvolutionDirection = __EvolutionDirection__00000007_
	__ParameterShape__00000000_.Parameter = __Parameter__00000002_
	__ParameterShape__00000001_.Parameter = __Parameter__00000003_
	__ParameterShape__00000002_.Parameter = __Parameter__00000001_
	__ParameterShape__00000003_.Parameter = __Parameter__00000004_
	__ParameterShape__00000004_.Parameter = __Parameter__00000005_
	__ParameterShape__00000006_.Parameter = __Parameter__00000000_
	__ParameterShape__00000007_.Parameter = __Parameter__00000006_
	__ParameterShape__00000008_.Parameter = __Parameter__00000006_
	__ParameterShape__00000009_.Parameter = __Parameter__00000007_
	__ParameterShape__00000010_.Parameter = __Parameter__00000008_
	__ParameterShape__00000012_.Parameter = __Parameter__00000009_
	__ParameterShape__00000013_.Parameter = __Parameter__00000010_
	__ParameterShape__00000014_.Parameter = __Parameter__00000011_
	__ParameterShape__00000015_.Parameter = __Parameter__00000012_
	__ParameterShape__00000016_.Parameter = __Parameter__00000013_
	__ParameterShape__00000017_.Parameter = __Parameter__00000014_
	__ParameterShape__00000018_.Parameter = __Parameter__00000015_
	__ParameterShape__00000019_.Parameter = __Parameter__00000016_
	__ParameterShape__00000020_.Parameter = __Parameter__00000017_
	__ParametersAggregate__00000001_.Parameters = append(__ParametersAggregate__00000001_.Parameters, __Parameter__00000004_)
	__ParametersAggregate__00000001_.Parameters = append(__ParametersAggregate__00000001_.Parameters, __Parameter__00000012_)
	__ParametersAggregate__00000001_.Parameters = append(__ParametersAggregate__00000001_.Parameters, __Parameter__00000015_)
	__ParametersAggregate__00000001_.Parameters = append(__ParametersAggregate__00000001_.Parameters, __Parameter__00000017_)
	__ParametersAggregate__00000001_.Parameters = append(__ParametersAggregate__00000001_.Parameters, __Parameter__00000016_)
	__ParametersAggregate__00000001_.Parameters = append(__ParametersAggregate__00000001_.Parameters, __Parameter__00000014_)
	__ParametersAggregate__00000002_.Parameters = append(__ParametersAggregate__00000002_.Parameters, __Parameter__00000002_)
	__ParametersAggregate__00000002_.Parameters = append(__ParametersAggregate__00000002_.Parameters, __Parameter__00000005_)
	__ParametersAggregate__00000002_.Parameters = append(__ParametersAggregate__00000002_.Parameters, __Parameter__00000007_)
	__ParametersAggregate__00000002_.Parameters = append(__ParametersAggregate__00000002_.Parameters, __Parameter__00000013_)
	__ParametersAggregate__00000003_.Parameters = append(__ParametersAggregate__00000003_.Parameters, __Parameter__00000001_)
	__ParametersAggregate__00000003_.Parameters = append(__ParametersAggregate__00000003_.Parameters, __Parameter__00000003_)
	__ParametersAggregate__00000003_.Parameters = append(__ParametersAggregate__00000003_.Parameters, __Parameter__00000008_)
	__ParametersAggregate__00000003_.Parameters = append(__ParametersAggregate__00000003_.Parameters, __Parameter__00000000_)
	__ParametersAggregate__00000004_.Parameters = append(__ParametersAggregate__00000004_.Parameters, __Parameter__00000010_)
	__ParametersAggregate__00000004_.Parameters = append(__ParametersAggregate__00000004_.Parameters, __Parameter__00000011_)
	__ParametersAggregateShape__00000000_.ScenarioParameter = __ParametersAggregate__00000004_
	__ParametersAggregateShape__00000001_.ScenarioParameter = __ParametersAggregate__00000001_
	__ParametersAggregateShape__00000003_.ScenarioParameter = __ParametersAggregate__00000002_
	__ParametersAggregateShape__00000004_.ScenarioParameter = __ParametersAggregate__00000003_
	__Scenario__00000000_.Diagrams = append(__Scenario__00000000_.Diagrams, __Diagram__00000001_)
	__Scenario__00000000_.ActorStates = append(__Scenario__00000000_.ActorStates, __ActorState__00000002_)
	__Scenario__00000000_.ActorStates = append(__Scenario__00000000_.ActorStates, __ActorState__00000003_)
	__Scenario__00000000_.ActorStates = append(__Scenario__00000000_.ActorStates, __ActorState__00000004_)
	__Scenario__00000000_.ActorStates = append(__Scenario__00000000_.ActorStates, __ActorState__00000005_)
	__Scenario__00000000_.ActorStateTransitions = append(__Scenario__00000000_.ActorStateTransitions, __ActorStateTransition__00000000_)
	__Scenario__00000000_.ActorStateTransitions = append(__Scenario__00000000_.ActorStateTransitions, __ActorStateTransition__00000003_)
	__Scenario__00000000_.ActorStateTransitions = append(__Scenario__00000000_.ActorStateTransitions, __ActorStateTransition__00000004_)
	__Scenario__00000000_.EvolutionDirections = append(__Scenario__00000000_.EvolutionDirections, __EvolutionDirection__00000002_)
	__Scenario__00000000_.EvolutionDirections = append(__Scenario__00000000_.EvolutionDirections, __EvolutionDirection__00000003_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000000_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000001_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000003_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000004_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000008_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000012_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000014_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000015_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000016_)
	__Scenario__00000000_.Parameters = append(__Scenario__00000000_.Parameters, __Parameter__00000017_)
	__Scenario__00000000_.ParametersAggretates = append(__Scenario__00000000_.ParametersAggretates, __ParametersAggregate__00000001_)
	__Scenario__00000000_.ParametersAggretates = append(__Scenario__00000000_.ParametersAggretates, __ParametersAggregate__00000003_)
	__Scenario__00000001_.Diagrams = append(__Scenario__00000001_.Diagrams, __Diagram__00000002_)
	__Scenario__00000001_.ActorStates = append(__Scenario__00000001_.ActorStates, __ActorState__00000000_)
	__Scenario__00000001_.ActorStates = append(__Scenario__00000001_.ActorStates, __ActorState__00000001_)
	__Scenario__00000001_.ActorStateTransitions = append(__Scenario__00000001_.ActorStateTransitions, __ActorStateTransition__00000001_)
	__Scenario__00000001_.EvolutionDirections = append(__Scenario__00000001_.EvolutionDirections, __EvolutionDirection__00000000_)
	__Scenario__00000001_.EvolutionDirections = append(__Scenario__00000001_.EvolutionDirections, __EvolutionDirection__00000007_)
	__Scenario__00000001_.Parameters = append(__Scenario__00000001_.Parameters, __Parameter__00000010_)
	__Scenario__00000001_.Parameters = append(__Scenario__00000001_.Parameters, __Parameter__00000011_)
	__Scenario__00000001_.ParametersAggretates = append(__Scenario__00000001_.ParametersAggretates, __ParametersAggregate__00000004_)
	__Scenario__00000002_.Diagrams = append(__Scenario__00000002_.Diagrams, __Diagram__00000003_)
	__Scenario__00000002_.ActorStates = append(__Scenario__00000002_.ActorStates, __ActorState__00000008_)
	__Scenario__00000002_.ActorStates = append(__Scenario__00000002_.ActorStates, __ActorState__00000009_)
	__Scenario__00000002_.ActorStateTransitions = append(__Scenario__00000002_.ActorStateTransitions, __ActorStateTransition__00000006_)
	__Scenario__00000002_.EvolutionDirections = append(__Scenario__00000002_.EvolutionDirections, __EvolutionDirection__00000005_)
	__Scenario__00000002_.EvolutionDirections = append(__Scenario__00000002_.EvolutionDirections, __EvolutionDirection__00000006_)
	__Scenario__00000002_.Parameters = append(__Scenario__00000002_.Parameters, __Parameter__00000002_)
	__Scenario__00000002_.Parameters = append(__Scenario__00000002_.Parameters, __Parameter__00000005_)
	__Scenario__00000002_.Parameters = append(__Scenario__00000002_.Parameters, __Parameter__00000007_)
	__Scenario__00000002_.Parameters = append(__Scenario__00000002_.Parameters, __Parameter__00000013_)
	__Scenario__00000002_.ParametersAggretates = append(__Scenario__00000002_.ParametersAggretates, __ParametersAggregate__00000000_)
	__Scenario__00000002_.ParametersAggretates = append(__Scenario__00000002_.ParametersAggretates, __ParametersAggregate__00000002_)
	__Scenario__00000005_.Diagrams = append(__Scenario__00000005_.Diagrams, __Diagram__00000000_)
	__Scenario__00000005_.Diagrams = append(__Scenario__00000005_.Diagrams, __Diagram__00000004_)
	__Scenario__00000005_.ActorStates = append(__Scenario__00000005_.ActorStates, __ActorState__00000006_)
	__Scenario__00000005_.ActorStates = append(__Scenario__00000005_.ActorStates, __ActorState__00000007_)
	__Scenario__00000005_.ActorStateTransitions = append(__Scenario__00000005_.ActorStateTransitions, __ActorStateTransition__00000005_)
	__Scenario__00000005_.EvolutionDirections = append(__Scenario__00000005_.EvolutionDirections, __EvolutionDirection__00000001_)
	__Scenario__00000005_.EvolutionDirections = append(__Scenario__00000005_.EvolutionDirections, __EvolutionDirection__00000004_)
	__Scenario__00000005_.Parameters = append(__Scenario__00000005_.Parameters, __Parameter__00000006_)
	__Scenario__00000005_.Parameters = append(__Scenario__00000005_.Parameters, __Parameter__00000009_)
	__Workspace__00000000_.SelectedDiagram = __Diagram__00000003_
	__Workspace__00000000_.Default_EvolutionDirectionShape = nil
	__Workspace__00000000_.Default_ParameterShape = nil
	__Workspace__00000000_.Default_ScenarioParameterShape = nil
	__Workspace__00000000_.Default_ActorStateShape = nil
	__Workspace__00000000_.Default_ActorStateTransitionShape = nil
}
