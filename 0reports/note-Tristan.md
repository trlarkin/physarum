Lab Notebook
============
This file hold all the information about results and how I achieved them.


24 April Findings
-----------------
Using the paper I determined settings that promote the development of a network connecting food sources. The important settings are:

```go
width:              1<<9
height:             1<<9
numParticles:       1<<17
blurRadius:         1
blurPasses:         2
zoomFactor:         1

SensorAngle:        Radians(45),
SensorDistance:     8,
RotationAngle:      Radians(45),
StepDistance:       1,
DepositionAmount:   2,
DecayFactor:        0.05,
```

### Summary
- I changed how food worked.
- I looked at changing which food maps I used and the number of particles. 
- I looked at the images normally, along with the log scaled versions. 
- I found that the number of particles could be a useful attribute to edit since more particles produces a larger network outside of the region with food points.
- The slimes seem to connect the closest food sources, but also have a tendency to create small world connections.
    - These connections are interesting since the small world connections change over time while the close connections pretty much always stay there. This might show how which small world connections are made doesn't really matter.
    - see this in `results/small-world-evolution/` 

### ... I will add more details later, but they are in my notebook, let me know if you want anything specific from the summary expanded on.
