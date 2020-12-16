
# TeamingTool

## Introduction 

TeamingTool is just a tool I built base on the idea of a personality test (https://www.truity.com/test/big-five-personality-test). There are 5 big personalities, and each of them theoretically can support another, as A support for A and C, O support for E and O, E support only for E, etc. You can relate this to the blood type for better consumption.

My work here required us to separated the class into group of 4 (or 3 or 5) that have the best performance based on work and self-development. However, doing the separation part by hand is quite s*ck, so I built this tool.

## Feature

This is just the server side that will be hosted by AWS. There will eventually be a UI for better usage.
The APIs takes the class info by Google-Sheet-API, which you will have to config in the main.go file. The Sheet column's structure should follow this:

Student | Name | Age | PA | PC | PE | PO | PN | Gender | 
--- | --- | --- | --- |--- |--- |--- |--- |--- |
Content | Phineas | 19 | 80 | 88 | 88 | 80 | 80 | Male | 
Content | Ferb | 19 | 80 | 88 | 88 | 80 | 80 | Male | 
Content | Isabella | 19 | 80 | 88 | 88 | 80 | 80 | Female | 

That's clear enough.

## Upgrade

This can be further implemented the truity test into tool. The truity test is will output only the array of `[]db.Student` that the `division.TeamDivision` need to make the teams. The truity test can be taken from the above links, and put anywhere but the `division` package.
