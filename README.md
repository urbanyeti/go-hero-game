# go-hero-game
An adventure simulator written in Go

## Packages

### character
Structs and helper functions for characters and objects in the game. Includes DTOs that can be inflated on game load.

### content
Contains tool for transforming game content (items, abilities, and characters) from Excel to JSON. When adding large amounts of data and tags, it's easier to manage everything in a big Excel workbook as the "source of truth". 

### grpc
Code for the protocal buffer, server, client to allow the content to be loaded and generated from an authenticated server.

### game
Core gameplay consisting of "loops" and "turns". With each turn of the adventure, the hero will acqure items, abilities, and bonuses or encounter combat. When combat starts, the combatants will automatically perform attacks based on their speed, damage, and other stats until only the victor remains. Loot and experience is dropped when the hero slays a monster.

#### Sample combat output
```
time="06-05-2021 21:49:42" level=info msg="turn started" game="Loop: 1 Turn: 1 | Dan - HP: 100 | Stats: {(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)(eva: 0)(agi: 5)(lvl: 1)(exp: 0)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]]"
time="06-05-2021 21:49:42" level=info msg="stat modified" character=monster-frog new=1 old=1 stat=lvl
time="06-05-2021 21:49:42" level=info msg="combat started" hero="Dan - HP: 100 | Stats: {(exp: 0)(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)(eva: 0)(agi: 5)(lvl: 1)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]]" monsters="map[monster-frog0:Giant Frog - HP: 8 | Stats: {(agi: 4)(atk: 6)(def: 6)(exp: 4)(lvl: 1)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[]]"
time="06-05-2021 21:49:42" level=warning msg="cannot retrieve missing stat" statID=eva
time="06-05-2021 21:49:42" level=info msg="damage taken" character=monster-frog dmg=3
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Dan - HP: 100 | Stats: {(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)(eva: 0)(agi: 5)(lvl: 1)(exp: 0)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] [Basic Sword] Giant Frog - HP: 5 | Stats: {(atk: 6)(def: 6)(exp: 4)(lvl: 1)(agi: 4)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] monster-frog0 84 80}" baseDmg=6 defenderDef=7 rollDmg=4 totalDmg=3
time="06-05-2021 21:49:42" level=warning msg="cannot retrieve missing stat" statID=eva
time="06-05-2021 21:49:42" level=info msg="damage taken" character=monster-frog dmg=1
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Dan - HP: 100 | Stats: {(atk: 5)(def: 5)(eva: 0)(agi: 5)(lvl: 1)(exp: 0)(exp-next: 10)(hp-max: 100)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] [Basic Sword] Giant Frog - HP: 4 | Stats: {(agi: 4)(atk: 6)(def: 6)(exp: 4)(lvl: 1)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] monster-frog0 84 80}" baseDmg=6 defenderDef=7 rollDmg=2 totalDmg=1
time="06-05-2021 21:49:42" level=info msg="damage taken" character=hero-dan dmg=9
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Giant Frog - HP: 8 | Stats: {(lvl: 1)(agi: 4)(atk: 6)(def: 6)(exp: 4)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] [Acid Spit] Dan - HP: 91 | Stats: {(agi: 5)(lvl: 1)(exp: 0)(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)(eva: 0)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] hero-dan 120 120}" baseDmg=7 defenderDef=9 rollDmg=11 totalDmg=9
time="06-05-2021 21:49:42" level=warning msg="cannot retrieve missing stat" statID=eva
time="06-05-2021 21:49:42" level=info msg="damage taken" character=monster-frog dmg=3
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Dan - HP: 91 | Stats: {(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)(eva: 0)(agi: 5)(lvl: 1)(exp: 0)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] [Basic Sword] Giant Frog - HP: 1 | Stats: {(agi: 4)(atk: 6)(def: 6)(exp: 4)(lvl: 1)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] monster-frog0 84 80}" baseDmg=6 defenderDef=7 rollDmg=4 totalDmg=3
time="06-05-2021 21:49:42" level=info msg="damage taken" character=hero-dan dmg=0
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Giant Frog - HP: 8 | Stats: {(agi: 4)(atk: 6)(def: 6)(exp: 4)(lvl: 1)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] [Tongue Smack] Dan - HP: 91 | Stats: {(eva: 0)(agi: 5)(lvl: 1)(exp: 0)(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] hero-dan 80 80}" baseDmg=7 defenderDef=9 rollDmg=1 totalDmg=0
time="06-05-2021 21:49:42" level=warning msg="cannot retrieve missing stat" statID=eva
time="06-05-2021 21:49:42" level=info msg="damage taken" character=monster-frog dmg=0
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Dan - HP: 91 | Stats: {(lvl: 1)(exp: 0)(exp-next: 10)(hp-max: 100)(atk: 5)(def: 5)(eva: 0)(agi: 5)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] [Basic Sword] Giant Frog - HP: 1 | Stats: {(def: 6)(exp: 4)(lvl: 1)(agi: 4)(atk: 6)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] monster-frog0 84 80}" baseDmg=6 defenderDef=7 rollDmg=1 totalDmg=0
time="06-05-2021 21:49:42" level=warning msg="cannot retrieve missing stat" statID=eva
time="06-05-2021 21:49:42" level=info msg="damage taken" character=monster-frog dmg=3
time="06-05-2021 21:49:42" level=info msg="damage dealt" attack="{Dan - HP: 91 | Stats: {(hp-max: 100)(atk: 5)(def: 5)(eva: 0)(agi: 5)(lvl: 1)(exp: 0)(exp-next: 10)} | map[] | map[arm-r:[Basic Sword] feet:[Basic Boots] torso:[Basic Leather Armor]] [Basic Sword] Giant Frog - HP: -2 | Stats: {(agi: 4)(atk: 6)(def: 6)(exp: 4)(lvl: 1)} | map[attack-spit1:[Acid Spit] attack-tongue1:[Tongue Smack]] | map[] monster-frog0 84 80}" baseDmg=6 defenderDef=7 rollDmg=4 totalDmg=3
time="06-05-2021 21:49:42" level=info msg="monster slain" hero=hero-dan monster=monster-frog
time="06-05-2021 21:49:42" level=info msg="stat modified" character=hero-dan new=4 old=0 stat=exp
time="06-05-2021 21:49:42" level=info msg="added item" character=hero-dan item=item-armor2
time="06-05-2021 21:49:42" level=info msg="combat finished" hero=hero-dan
```



