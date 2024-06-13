Note: the code structure might resemble Java, as I have never worked with Go before. This code is written in Go because it provides a simple concurrency model.

Overview:
This code implements a small problem of booking systems: seat booking. 

How does it work?
- There are open sockets through which a number of users can ask for seat booking.
- Seats are booked in another thread and it each booking thread launches a goroutine to get the payment, if payment is done within 5 seconds, seat is booked, otherwise released. Also, the paymet service is just a dummy service.
- Seats and the whole seating arrangement is managed using an sql db.

### Todos:
- [ ] Run a docker sql container : create a script to bootstrap docker.
- [x] Create connection to sql
- [x] For a v0 of the project, store seat and seating info in a sql db table
- [ ] create the basic flow: a user starts the booking flow and the system waits 5 seconds for the payment, if payment is not done, the abort booking
- [ ] create a load tester script: 100 users/second booking