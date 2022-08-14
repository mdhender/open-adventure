// Code generated by yaml2go on 2022-08-14T18:50:47Z; DO NOT EDIT.

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package state

// Objects is an ordered map of objects
// Use `Seq` to access sequentially or `Map` to access by key.
type Objects struct {
    Seq []*Object            // sequential accessor
    Map map[string]*Object   // tag accessor
}

// Object is autogenerated
type Object struct {
    Changes      []string
    Descriptions []string
    Locations    []string
    Immovable    bool
    Inventory    string
    States       []string
    Treasure     bool
    Words        []string
}

// generateObjects returns the initial state of objects
func generateObjects() *Objects {
    objects := &Objects{
        Map: make(map[string]*Object),
    }

    // seq: 0 name: NO_OBJECT
    object := &Object{
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["NO_OBJECT"] = object

    // seq: 1 name: KEYS
    object = &Object{
        Descriptions: []string{
            "There are some keys on the ground here.",
        },
        Inventory: "Set of keys",
        Locations: []string{
            "LOC_BUILDING",
        },
        Words: []string{
            "keys",
            "key",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["KEYS"] = object

    // seq: 2 name: LAMP
    object = &Object{
        Changes: []string{
            "Your lamp is now off.",
            "Your lamp is now on.",
        },
        Descriptions: []string{
            "There is a shiny brass lamp nearby.",
            "There is a lamp shining nearby.",
        },
        Inventory: "Brass lantern",
        Locations: []string{
            "LOC_BUILDING",
        },
        States: []string{
            "LAMP_DARK",
            "LAMP_BRIGHT",
        },
        Words: []string{
            "lamp",
            "lante",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["LAMP"] = object

    // seq: 3 name: GRATE
    object = &Object{
        Changes: []string{
            "The grate is now locked.",
            "The grate is now unlocked.",
        },
        Descriptions: []string{
            "The grate is locked.",
            "The grate is open.",
        },
        Immovable: true,
        Inventory: "*grate",
        Locations: []string{
            "LOC_GRATE",
            "LOC_BELOWGRATE",
        },
        States: []string{
            "GRATE_CLOSED",
            "GRATE_OPEN",
        },
        Words: []string{
            "grate",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["GRATE"] = object

    // seq: 4 name: CAGE
    object = &Object{
        Descriptions: []string{
            "There is a small wicker cage discarded nearby.",
        },
        Inventory: "Wicker cage",
        Locations: []string{
            "LOC_COBBLE",
        },
        Words: []string{
            "cage",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["CAGE"] = object

    // seq: 5 name: ROD
    object = &Object{
        Descriptions: []string{
            "A three foot black rod with a rusty star on an end lies nearby.",
        },
        Inventory: "Black rod",
        Locations: []string{
            "LOC_DEBRIS",
        },
        Words: []string{
            "rod",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["ROD"] = object

    // seq: 6 name: ROD2
    object = &Object{
        Descriptions: []string{
            "A three foot black rod with a rusty mark on an end lies nearby.",
        },
        Inventory: "Black rod",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "rod",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["ROD2"] = object

    // seq: 7 name: STEPS
    object = &Object{
        Descriptions: []string{
            "Rough stone steps lead down the pit.",
            "Rough stone steps lead up the dome.",
        },
        Immovable: true,
        Inventory: "*steps",
        Locations: []string{
            "LOC_PITTOP",
            "LOC_MISTHALL",
        },
        States: []string{
            "STEPS_DOWN",
            "STEPS_UP",
        },
        Words: []string{
            "steps",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["STEPS"] = object

    // seq: 8 name: BIRD
    object = &Object{
        Descriptions: []string{
            "A cheerful little bird is sitting here singing.",
            "There is a little bird in the cage.",
            "A cheerful little bird is sitting here singing.",
        },
        Inventory: "Little bird in cage",
        Locations: []string{
            "LOC_BIRDCHAMBER",
        },
        States: []string{
            "BIRD_UNCAGED",
            "BIRD_CAGED",
            "BIRD_FOREST_UNCAGED",
        },
        Words: []string{
            "bird",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["BIRD"] = object

    // seq: 9 name: DOOR
    object = &Object{
        Changes: []string{
            "The hinges are quite thoroughly rusted now and won't budge.",
            "The oil has freed up the hinges so that the door will now move,\nalthough it requires some effort.",
        },
        Descriptions: []string{
            "The way north is barred by a massive, rusty, iron door.",
            "The way north leads through a massive, rusty, iron door.",
        },
        Immovable: true,
        Inventory: "*rusty door",
        Locations: []string{
            "LOC_IMMENSE",
        },
        States: []string{
            "DOOR_RUSTED",
            "DOOR_UNRUSTED",
        },
        Words: []string{
            "door",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["DOOR"] = object

    // seq: 10 name: PILLOW
    object = &Object{
        Descriptions: []string{
            "A small velvet pillow lies on the floor.",
        },
        Inventory: "Velvet pillow",
        Locations: []string{
            "LOC_SOFTROOM",
        },
        Words: []string{
            "pillo",
            "velve",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["PILLOW"] = object

    // seq: 11 name: SNAKE
    object = &Object{
        Descriptions: []string{
            "A huge green fierce snake bars the way!",
            "",
        },
        Immovable: true,
        Inventory: "*snake",
        Locations: []string{
            "LOC_KINGHALL",
        },
        States: []string{
            "SNAKE_BLOCKS",
            "SNAKE_CHASED",
        },
        Words: []string{
            "snake",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["SNAKE"] = object

    // seq: 12 name: FISSURE
    object = &Object{
        Changes: []string{
            "The crystal bridge has vanished!",
            "A crystal bridge now spans the fissure.",
        },
        Descriptions: []string{
            "",
            "A crystal bridge spans the fissure.",
        },
        Immovable: true,
        Inventory: "*fissure",
        Locations: []string{
            "LOC_EASTBANK",
            "LOC_WESTBANK",
        },
        States: []string{
            "UNBRIDGED",
            "BRIDGED",
        },
        Words: []string{
            "fissu",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["FISSURE"] = object

    // seq: 13 name: OBJ_13
    object = &Object{
        Descriptions: []string{
            "A massive stone tablet embedded in the wall reads:\n\"Congratulations on bringing light into the dark-room!\"",
        },
        Immovable: true,
        Inventory: "*stone tablet",
        Locations: []string{
            "LOC_DARKROOM",
        },
        Words: []string{
            "table",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_13"] = object

    // seq: 14 name: CLAM
    object = &Object{
        Descriptions: []string{
            "There is an enormous clam here with its shell tightly closed.",
        },
        Inventory: "Giant clam  >GRUNT!<",
        Locations: []string{
            "LOC_SHELLROOM",
        },
        Words: []string{
            "clam",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["CLAM"] = object

    // seq: 15 name: OYSTER
    object = &Object{
        Descriptions: []string{
            "There is an enormous oyster here with its shell tightly closed.",
            "Interesting.  There seems to be something written on the underside of\\nthe oyster.",
        },
        Inventory: "Giant oyster  >GROAN!<",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "oyste",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OYSTER"] = object

    // seq: 16 name: MAGAZINE
    object = &Object{
        Descriptions: []string{
            "There are a few recent issues of \"Spelunker Today\" magazine here.",
        },
        Inventory: "\"Spelunker Today\"",
        Locations: []string{
            "LOC_ANTEROOM",
        },
        Words: []string{
            "magaz",
            "issue",
            "spelu",
            "\"spel",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["MAGAZINE"] = object

    // seq: 17 name: DWARF
    object = &Object{
        Immovable: true,
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "dwarf",
            "dwarv",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["DWARF"] = object

    // seq: 18 name: KNIFE
    object = &Object{
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "knife",
            "knive",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["KNIFE"] = object

    // seq: 19 name: FOOD
    object = &Object{
        Descriptions: []string{
            "There is food here.",
        },
        Inventory: "Tasty food",
        Locations: []string{
            "LOC_BUILDING",
        },
        Words: []string{
            "food",
            "ratio",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["FOOD"] = object

    // seq: 20 name: BOTTLE
    object = &Object{
        Changes: []string{
            "Your bottle is now full of water.",
            "The bottle of water is now empty.",
            "Your bottle is now full of oil.",
        },
        Descriptions: []string{
            "There is a bottle of water here.",
            "There is an empty bottle here.",
            "There is a bottle of oil here.",
        },
        Inventory: "Small bottle",
        Locations: []string{
            "LOC_BUILDING",
        },
        States: []string{
            "WATER_BOTTLE",
            "EMPTY_BOTTLE",
            "OIL_BOTTLE",
        },
        Words: []string{
            "bottl",
            "jar",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["BOTTLE"] = object

    // seq: 21 name: WATER
    object = &Object{
        Inventory: "Water in the bottle",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "water",
            "h2o",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["WATER"] = object

    // seq: 22 name: OIL
    object = &Object{
        Inventory: "Oil in the bottle",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "oil",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OIL"] = object

    // seq: 23 name: MIRROR
    object = &Object{
        Changes: []string{
            "",
            "You strike the mirror a resounding blow, whereupon it shatters into a\nmyriad tiny fragments.",
        },
        Descriptions: []string{
            "",
            "",
        },
        Immovable: true,
        Inventory: "*mirror",
        Locations: []string{
            "LOC_MIRRORCANYON",
        },
        States: []string{
            "MIRROR_UNBROKEN",
            "MIRROR_BROKEN",
        },
        Words: []string{
            "mirro",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["MIRROR"] = object

    // seq: 24 name: PLANT
    object = &Object{
        Changes: []string{
            "You've over-watered the plant!  It's shriveling up!  And now . . .",
            "The plant spurts into furious growth for a few seconds.",
            "The plant grows explosively, almost filling the bottom of the pit.",
        },
        Descriptions: []string{
            "There is a tiny little plant in the pit, murmuring \"water, water, ...\"",
            "There is a 12-foot-tall beanstalk stretching up out of the pit,\\nbellowing \"WATER!! WATER!!\"",
            "There is a gigantic beanstalk stretching all the way up to the hole.",
        },
        Immovable: true,
        Inventory: "*plant",
        Locations: []string{
            "LOC_WESTPIT",
        },
        States: []string{
            "PLANT_THIRSTY",
            "PLANT_BELLOWING",
            "PLANT_GROWN",
        },
        Words: []string{
            "plant",
            "beans",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["PLANT"] = object

    // seq: 25 name: PLANT2
    object = &Object{
        Descriptions: []string{
            "",
            "The top of a 12-foot-tall beanstalk is poking out of the west pit.",
            "There is a huge beanstalk growing out of the west pit up to the hole.",
        },
        Immovable: true,
        Inventory: "*phony plant",
        Locations: []string{
            "LOC_WESTEND",
            "LOC_EASTEND",
        },
        Words: []string{
            "plant",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["PLANT2"] = object

    // seq: 26 name: OBJ_26
    object = &Object{
        Descriptions: []string{
            "",
        },
        Immovable: true,
        Inventory: "*stalactite",
        Locations: []string{
            "LOC_TOPSTALACTITE",
        },
        Words: []string{
            "stala",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_26"] = object

    // seq: 27 name: OBJ_27
    object = &Object{
        Descriptions: []string{
            "The shadowy figure seems to be trying to attract your attention.",
        },
        Immovable: true,
        Inventory: "*shadowy figure and/or window",
        Locations: []string{
            "LOC_WINDOW1",
            "LOC_WINDOW2",
        },
        Words: []string{
            "shado",
            "figur",
            "windo",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_27"] = object

    // seq: 28 name: AXE
    object = &Object{
        Changes: []string{
            "",
            "The axe misses and lands near the bear where you can't get at it.",
        },
        Descriptions: []string{
            "There is a little axe here.",
            "There is a little axe lying beside the bear.",
        },
        Inventory: "Dwarf's axe",
        Locations: []string{
            "LOC_NOWHERE",
        },
        States: []string{
            "AXE_HERE",
            "AXE_LOST",
        },
        Words: []string{
            "axe",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["AXE"] = object

    // seq: 29 name: OBJ_29
    object = &Object{
        Immovable: true,
        Inventory: "*cave drawings",
        Locations: []string{
            "LOC_ORIENTAL",
        },
        Words: []string{
            "drawi",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_29"] = object

    // seq: 30 name: OBJ_30
    object = &Object{
        Immovable: true,
        Inventory: "*pirate/genie",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "pirat",
            "genie",
            "djinn",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_30"] = object

    // seq: 31 name: DRAGON
    object = &Object{
        Changes: []string{
            "",
            "Congratulations!  You have just vanquished a dragon with your bare\nhands!  (Unbelievable, isn't it?)",
            "Your head buzzes strangely for a moment.",
        },
        Descriptions: []string{
            "A huge green fierce dragon bars the way!",
            "The blood-specked body of a huge green dead dragon lies to one side.",
            "The body of a huge green dead dragon is lying off to one side.",
        },
        Immovable: true,
        Inventory: "*dragon",
        Locations: []string{
            "LOC_SECRET4",
            "LOC_SECRET6",
        },
        States: []string{
            "DRAGON_BARS",
            "DRAGON_DEAD",
            "DRAGON_BLOODLESS",
        },
        Words: []string{
            "drago",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["DRAGON"] = object

    // seq: 32 name: CHASM
    object = &Object{
        Changes: []string{
            "",
            "Just as you reach the other side, the bridge buckles beneath the\nweight of the bear, which was still following you around.  You\nscrabble desperately for support, but as the bridge collapses you\nstumble back and fall into the chasm.",
        },
        Descriptions: []string{
            "A rickety wooden bridge extends across the chasm, vanishing into the\\nmist.  A notice posted on the bridge reads, \"Stop! Pay troll!\"",
            "The wreckage of a bridge (and a dead bear) can be seen at the bottom\\nof the chasm.",
        },
        Immovable: true,
        Inventory: "*chasm",
        Locations: []string{
            "LOC_SWCHASM",
            "LOC_NECHASM",
        },
        States: []string{
            "TROLL_BRIDGE",
            "BRIDGE_WRECKED",
        },
        Words: []string{
            "chasm",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["CHASM"] = object

    // seq: 33 name: TROLL
    object = &Object{
        Changes: []string{
            "",
            "",
            "The bear lumbers toward the troll, who lets out a startled shriek and\nscurries away.  The bear soon gives up the pursuit and wanders back.",
        },
        Descriptions: []string{
            "A burly troll stands by the bridge and insists you throw him a\\ntreasure before you may cross.",
            "The troll steps out from beneath the bridge and blocks your way.",
            "",
        },
        Immovable: true,
        Inventory: "*troll",
        Locations: []string{
            "LOC_SWCHASM",
            "LOC_NECHASM",
        },
        States: []string{
            "TROLL_UNPAID",
            "TROLL_PAIDONCE",
            "TROLL_GONE",
        },
        Words: []string{
            "troll",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["TROLL"] = object

    // seq: 34 name: TROLL2
    object = &Object{
        Descriptions: []string{
            "The troll is nowhere to be seen.",
        },
        Immovable: true,
        Inventory: "*phony troll",
        Locations: []string{
            "LOC_NOWHERE",
            "LOC_NOWHERE",
        },
        Words: []string{
            "troll",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["TROLL2"] = object

    // seq: 35 name: BEAR
    object = &Object{
        Changes: []string{
            "",
            "The bear eagerly wolfs down your food, after which he seems to calm\\ndown considerably and even becomes rather friendly.",
            "",
            "",
        },
        Descriptions: []string{
            "There is a ferocious cave bear eyeing you from the far end of the room!",
            "There is a gentle cave bear sitting placidly in one corner.",
            "There is a contented-looking bear wandering about nearby.",
            "",
        },
        Immovable: true,
        Locations: []string{
            "LOC_BARRENROOM",
        },
        States: []string{
            "UNTAMED_BEAR",
            "SITTING_BEAR",
            "CONTENTED_BEAR",
            "BEAR_DEAD",
        },
        Words: []string{
            "bear",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["BEAR"] = object

    // seq: 36 name: MESSAG
    object = &Object{
        Descriptions: []string{
            "There is a message scrawled in the dust in a flowery script, reading:\n\"This is not the maze where the pirate leaves his treasure chest.\"",
        },
        Immovable: true,
        Inventory: "*message in second maze",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "messa",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["MESSAG"] = object

    // seq: 37 name: VOLCANO
    object = &Object{
        Immovable: true,
        Inventory: "*volcano and/or geyser",
        Locations: []string{
            "LOC_BREATHTAKING",
        },
        Words: []string{
            "volca",
            "geyse",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["VOLCANO"] = object

    // seq: 38 name: VEND
    object = &Object{
        Changes: []string{
            "The vending machine swings back to block the passage.",
            "As you strike the vending machine, it pivots backward along with a\\nsection of wall, revealing a dark passage leading south.",
        },
        Descriptions: []string{
            "There is a massive and somewhat battered vending machine here.  The\ninstructions on it read: \"Drop coins here to receive fresh batteries.\"",
            "There is a massive vending machine here, swung back to reveal a\nsouthward passage.",
        },
        Immovable: true,
        Inventory: "*vending machine",
        Locations: []string{
            "LOC_DEADEND13",
        },
        States: []string{
            "VEND_BLOCKS",
            "VEND_UNBLOCKS",
        },
        Words: []string{
            "machi",
            "vendi",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["VEND"] = object

    // seq: 39 name: BATTERY
    object = &Object{
        Descriptions: []string{
            "There are fresh batteries here.",
            "Some worn-out batteries have been discarded nearby.",
        },
        Inventory: "Batteries",
        Locations: []string{
            "LOC_NOWHERE",
        },
        States: []string{
            "FRESH_BATTERIES",
            "DEAD_BATTERIES",
        },
        Words: []string{
            "batte",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["BATTERY"] = object

    // seq: 40 name: OBJ_40
    object = &Object{
        Immovable: true,
        Inventory: "*carpet and/or moss and/or curtains",
        Locations: []string{
            "LOC_SOFTROOM",
        },
        Words: []string{
            "carpe",
            "moss",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_40"] = object

    // seq: 41 name: OGRE
    object = &Object{
        Descriptions: []string{
            "A formidable ogre bars the northern exit.",
        },
        Immovable: true,
        Inventory: "*ogre",
        Locations: []string{
            "LOC_LARGE",
        },
        Words: []string{
            "ogre",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OGRE"] = object

    // seq: 42 name: URN
    object = &Object{
        Changes: []string{
            "The urn is empty and will not light.",
            "The urn is now dark.",
            "The urn is now lit.",
        },
        Descriptions: []string{
            "A small urn is embedded in the rock.",
            "A small urn full of oil is embedded in the rock.",
            "A small oil flame extrudes from an urn embedded in the rock.",
        },
        Immovable: true,
        Inventory: "*urn",
        Locations: []string{
            "LOC_CLIFF",
        },
        States: []string{
            "URN_EMPTY",
            "URN_DARK",
            "URN_LIT",
        },
        Words: []string{
            "urn",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["URN"] = object

    // seq: 43 name: CAVITY
    object = &Object{
        Descriptions: []string{
            "",
            "There is a small urn-shaped cavity in the rock.",
        },
        Immovable: true,
        Inventory: "*cavity",
        Locations: []string{
            "LOC_NOWHERE",
        },
        States: []string{
            "CAVITY_FULL",
            "CAVITY_EMPTY",
        },
        Words: []string{
            "cavit",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["CAVITY"] = object

    // seq: 44 name: BLOOD
    object = &Object{
        Descriptions: []string{
            "",
        },
        Immovable: true,
        Inventory: "*blood",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Words: []string{
            "blood",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["BLOOD"] = object

    // seq: 45 name: RESER
    object = &Object{
        Changes: []string{
            "The waters crash together again.",
            "The waters have parted to form a narrow path across the reservoir.",
        },
        Descriptions: []string{
            "",
            "The waters have parted to form a narrow path across the reservoir.",
        },
        Immovable: true,
        Inventory: "*reservoir",
        Locations: []string{
            "LOC_RESERVOIR",
            "LOC_RESNORTH",
        },
        States: []string{
            "WATERS_UNPARTED",
            "WATERS_PARTED",
        },
        Words: []string{
            "reser",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["RESER"] = object

    // seq: 46 name: RABBITFOOT
    object = &Object{
        Descriptions: []string{
            "Your keen eye spots a severed leporine appendage lying on the ground.",
        },
        Inventory: "Leporine appendage",
        Locations: []string{
            "LOC_FOREST22",
        },
        Words: []string{
            "appen",
            "lepor",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["RABBITFOOT"] = object

    // seq: 47 name: OBJ_47
    object = &Object{
        Descriptions: []string{
            "",
        },
        Immovable: true,
        Inventory: "*mud",
        Locations: []string{
            "LOC_DEBRIS",
        },
        Words: []string{
            "mud",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_47"] = object

    // seq: 48 name: OBJ_48
    object = &Object{
        Descriptions: []string{
            "",
        },
        Immovable: true,
        Inventory: "*note",
        Locations: []string{
            "LOC_NUGGET",
        },
        Words: []string{
            "note",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_48"] = object

    // seq: 49 name: SIGN
    object = &Object{
        Descriptions: []string{
            "",
            "",
        },
        Immovable: true,
        Inventory: "*sign",
        Locations: []string{
            "LOC_ANTEROOM",
        },
        States: []string{
            "INGAME_SIGN",
            "ENDGAME_SIGN",
        },
        Words: []string{
            "sign",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["SIGN"] = object

    // seq: 50 name: NUGGET
    object = &Object{
        Descriptions: []string{
            "There is a large sparkling nugget of gold here!",
        },
        Inventory: "Large gold nugget",
        Locations: []string{
            "LOC_NUGGET",
        },
        Treasure: true,
        Words: []string{
            "gold",
            "nugge",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["NUGGET"] = object

    // seq: 51 name: OBJ_51
    object = &Object{
        Descriptions: []string{
            "There are diamonds here!",
        },
        Inventory: "Several diamonds",
        Locations: []string{
            "LOC_WESTBANK",
        },
        Treasure: true,
        Words: []string{
            "diamo",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_51"] = object

    // seq: 52 name: OBJ_52
    object = &Object{
        Descriptions: []string{
            "There are bars of silver here!",
        },
        Inventory: "Bars of silver",
        Locations: []string{
            "LOC_FLOORHOLE",
        },
        Treasure: true,
        Words: []string{
            "silve",
            "bars",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_52"] = object

    // seq: 53 name: OBJ_53
    object = &Object{
        Descriptions: []string{
            "There is precious jewelry here!",
        },
        Inventory: "Precious jewelry",
        Locations: []string{
            "LOC_SOUTHSIDE",
        },
        Treasure: true,
        Words: []string{
            "jewel",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_53"] = object

    // seq: 54 name: COINS
    object = &Object{
        Descriptions: []string{
            "There are many coins here!",
        },
        Inventory: "Rare coins",
        Locations: []string{
            "LOC_WESTSIDE",
        },
        Treasure: true,
        Words: []string{
            "coins",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["COINS"] = object

    // seq: 55 name: CHEST
    object = &Object{
        Descriptions: []string{
            "The pirate's treasure chest is here!",
        },
        Inventory: "Treasure chest",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Treasure: true,
        Words: []string{
            "chest",
            "box",
            "treas",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["CHEST"] = object

    // seq: 56 name: EGGS
    object = &Object{
        Descriptions: []string{
            "There is a large nest here, full of golden eggs!",
            "The nest of golden eggs has vanished!",
            "Done!",
        },
        Inventory: "Golden eggs",
        Locations: []string{
            "LOC_GIANTROOM",
        },
        States: []string{
            "EGGS_HERE",
            "EGGS_VANISHED",
            "EGGS_DONE",
        },
        Treasure: true,
        Words: []string{
            "eggs",
            "egg",
            "nest",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["EGGS"] = object

    // seq: 57 name: TRIDENT
    object = &Object{
        Descriptions: []string{
            "There is a jewel-encrusted trident here!",
        },
        Inventory: "Jeweled trident",
        Locations: []string{
            "LOC_WATERFALL",
        },
        Treasure: true,
        Words: []string{
            "tride",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["TRIDENT"] = object

    // seq: 58 name: VASE
    object = &Object{
        Changes: []string{
            "The vase is now resting, delicately, on a velvet pillow.",
            "The ming vase drops with a delicate crash.",
            "You have taken the vase and hurled it delicately to the ground.",
        },
        Descriptions: []string{
            "There is a delicate, precious, ming vase here!",
            "The floor is littered with worthless shards of pottery.",
            "The floor is littered with worthless shards of pottery.",
        },
        Inventory: "Ming vase",
        Locations: []string{
            "LOC_ORIENTAL",
        },
        States: []string{
            "VASE_WHOLE",
            "VASE_DROPPED",
            "VASE_BROKEN",
        },
        Treasure: true,
        Words: []string{
            "vase",
            "ming",
            "shard",
            "potte",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["VASE"] = object

    // seq: 59 name: EMERALD
    object = &Object{
        Descriptions: []string{
            "There is an emerald here the size of a plover's egg!",
            "There is an emerald resting in a small cavity in the rock!",
        },
        Inventory: "Egg-sized emerald",
        Locations: []string{
            "LOC_PLOVER",
        },
        Treasure: true,
        Words: []string{
            "emera",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["EMERALD"] = object

    // seq: 60 name: PYRAMID
    object = &Object{
        Descriptions: []string{
            "There is a platinum pyramid here, 8 inches on a side!",
        },
        Inventory: "Platinum pyramid",
        Locations: []string{
            "LOC_DARKROOM",
        },
        Treasure: true,
        Words: []string{
            "plati",
            "pyram",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["PYRAMID"] = object

    // seq: 61 name: PEARL
    object = &Object{
        Descriptions: []string{
            "Off to one side lies a glistening pearl!",
        },
        Inventory: "Glistening pearl",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Treasure: true,
        Words: []string{
            "pearl",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["PEARL"] = object

    // seq: 62 name: RUG
    object = &Object{
        Descriptions: []string{
            "There is a Persian rug spread out on the floor!",
            "The dragon is sprawled out on a Persian rug!!",
            "There is a Persian rug here, hovering in mid-air!",
        },
        Immovable: true,
        Inventory: "Persian rug",
        Locations: []string{
            "LOC_SECRET4",
            "LOC_SECRET6",
        },
        States: []string{
            "RUG_FLOOR",
            "RUG_DRAGON",
            "RUG_HOVER",
        },
        Treasure: true,
        Words: []string{
            "rug",
            "persi",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["RUG"] = object

    // seq: 63 name: OBJ_63
    object = &Object{
        Descriptions: []string{
            "There are rare spices here!",
        },
        Inventory: "Rare spices",
        Locations: []string{
            "LOC_BOULDERS2",
        },
        Treasure: true,
        Words: []string{
            "spice",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_63"] = object

    // seq: 64 name: CHAIN
    object = &Object{
        Descriptions: []string{
            "There is a golden chain lying in a heap on the floor!",
            "The bear is locked to the wall with a golden chain!",
            "There is a golden chain locked to the wall!",
        },
        Immovable: true,
        Inventory: "Golden chain",
        Locations: []string{
            "LOC_BARRENROOM",
        },
        States: []string{
            "CHAIN_HEAP",
            "CHAINING_BEAR",
            "CHAIN_FIXED",
        },
        Treasure: true,
        Words: []string{
            "chain",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["CHAIN"] = object

    // seq: 65 name: RUBY
    object = &Object{
        Descriptions: []string{
            "There is an enormous ruby here!",
            "There is a ruby resting in a small cavity in the rock!",
        },
        Inventory: "Giant ruby",
        Locations: []string{
            "LOC_STOREROOM",
        },
        Treasure: true,
        Words: []string{
            "ruby",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["RUBY"] = object

    // seq: 66 name: JADE
    object = &Object{
        Descriptions: []string{
            "A precious jade necklace has been dropped here!",
        },
        Inventory: "Jade necklace",
        Locations: []string{
            "LOC_NOWHERE",
        },
        Treasure: true,
        Words: []string{
            "jade",
            "neckl",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["JADE"] = object

    // seq: 67 name: AMBER
    object = &Object{
        Descriptions: []string{
            "There is a rare amber gemstone here!",
            "There is an amber gemstone resting in a small cavity in the rock!",
        },
        Inventory: "Amber gemstone",
        Locations: []string{
            "LOC_NOWHERE",
        },
        States: []string{
            "AMBER_IN_URN",
            "AMBER_IN_ROCK",
        },
        Treasure: true,
        Words: []string{
            "amber",
            "gemst",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["AMBER"] = object

    // seq: 68 name: SAPPH
    object = &Object{
        Descriptions: []string{
            "A brilliant blue star sapphire is here!",
            "There is a star sapphire resting in a small cavity in the rock!",
        },
        Inventory: "Star sapphire",
        Locations: []string{
            "LOC_LEDGE",
        },
        Treasure: true,
        Words: []string{
            "sapph",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["SAPPH"] = object

    // seq: 69 name: OBJ_69
    object = &Object{
        Descriptions: []string{
            "There is a richly-carved ebony statuette here!",
        },
        Inventory: "Ebony statuette",
        Locations: []string{
            "LOC_REACHDEAD",
        },
        Treasure: true,
        Words: []string{
            "ebony",
            "statu",
        },
    }
    objects.Seq = append(objects.Seq, object)
    objects.Map["OBJ_69"] = object

    return objects
}
