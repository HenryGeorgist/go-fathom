# go-fathom

A package to leverage go-consequences with fathom data

This library provides tools to support evaluating natural hazards interacting with concequence receptors. An example would be a flood represented by depth interacting with a residential structure to produce an estimate of economic losses at the residential structure.

## Packages
The following packages are in this library:
- compute
- hazard_providers
- store

### compute
The compute package computes damages for each frequency based event stored in the fathom data, and summarizes it to EAD per structure and writes to database.

### hazard_providers
The hazard providers package is designed to fulfil the interface HazardProvider with the fathom data.

### store
The store package is designed to interoperate with a local sqlite store to provide storage of results from national scale computes.

## Testing
Tests have been developed for most of the code related to flood damage estimation. The tests can be compiled using the general calls listed below on a package level. 

```
C:\Examples\Go_Consequences>go test ./paireddata -c
C:\Examples\Go_Consequences>.\paireddata.test -test.v