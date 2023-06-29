## I. Why use concurrency?
#### 1. Parallel:
- Parallel means 2 things at once, at the same time
- Advantage: complete task faster
- Some tasks must be performed sequentially

Remember: **some tasks are parallelizable and some aren't!**

#### 2. Von Neumann Bottlenecks
* Speed up without Parallelism
    
    * Faster processors
    * Cache

* We can't rely on hardware forever

#### 3. Power wall
- Increase transistor density -> Increase power consumption
- High power = High temperature

- Solution: multicore system? - still increase density

=> Parallel execution is needed to explout multi-core systems

Concurrent Programming: utilize multi-core system to develop software

#### 4. Concurrent vs Parallel
- Concurrent: start and end times overlap
- Parallel: execute at exactly the same time
- Programmer determines tasks that can be executed in parallel
- Mapping tasks to hardware:
    - OS
    - Go scheduler