// Code generated by yaml2go on 2022-08-17T04:50:23Z; DO NOT EDIT.

// Open Adventure - a clone of Open Adventure written in Go
// Copyright (c) 2022 Michael D Henderson
// SPDX-License-Identifier: BSD-2-clause

package state

// Class holds data for rating by points scored.
type Class struct {
    Threshold int       // minimum score needed for the rating
    Message   []string
}

// generateClasses returns the initial state for classes
func generateClasses() []*Class {
    return []*Class{
        &Class{
        },
        &Class{
            Threshold: 45,
            Message: []string{
                "You are obviously a rank amateur.  Better luck next time.",
            },
        },
        &Class{
            Threshold: 120,
            Message: []string{
                "Your score qualifies you as a novice class adventurer.",
            },
        },
        &Class{
            Threshold: 170,
            Message: []string{
                "You have achieved the rating: \"Experienced Adventurer\".",
            },
        },
        &Class{
            Threshold: 250,
            Message: []string{
                "You may now consider yourself a \"Seasoned Adventurer\".",
            },
        },
        &Class{
            Threshold: 320,
            Message: []string{
                "You have reached \"Junior Master\" status.",
            },
        },
        &Class{
            Threshold: 375,
            Message: []string{
                "Your score puts you in Master Adventurer Class C.",
            },
        },
        &Class{
            Threshold: 410,
            Message: []string{
                "Your score puts you in Master Adventurer Class B.",
            },
        },
        &Class{
            Threshold: 426,
            Message: []string{
                "Your score puts you in Master Adventurer Class A.",
            },
        },
        &Class{
            Threshold: 429,
            Message: []string{
                "All of Adventuredom gives tribute to you, Adventurer Grandmaster!",
            },
        },
        &Class{
            Threshold: 9999,
            Message: []string{
                "'Adventuredom stands in awe -- you have now joined the ranks of the",
                "       W O R L D   C H A M P I O N   A D V E N T U R E R S !",
                "It may interest you to know that the Dungeon-Master himself has, to",
                "my knowledge, never achieved this threshold in fewer than 330 turns.'",
            },
        },
    }
}

