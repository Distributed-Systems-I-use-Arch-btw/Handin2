# Handin2

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

    We send slices of bytes.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

    Our implementation uses theads. Thrads are unrealistic since threads have shared memory space, while on the other hand processes run on separate memory space.
    And in the real world client and servers dont have shared memeory.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

    We would assign each segment a sequence number. So that the client knows which order theyre sent no matter the order its recieved. 

d) In case messages can be delayed or lost, how does your implementation handle message loss?

    Each segment has a timeout, if there isnt recieved an acknolegdement in a certian timespan, we timeout and resend the semgnet.

e) Why is the 3-way handshake important?

    It's great to have a connection, where both the client and server have acknoledged each other which ensures clear communication without the loss of data.