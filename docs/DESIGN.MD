# EasyPark design

## Parking Request

### States

1. A parking request can be assigned a parking space if:
        - request's time slot doesn't overlap with existing parking requests for that parking space.
        - request's start time is not in the past.
        - request is not rejected.
        - space status is not 'blocked'.
        - space is in the desired parking lot.

## Parking Space

### States

1. Creation
    - Upon creation, all parking spaces are 'available'.

2. Parking request is assigned a parking space
    - When a driver's parking request is assigned a parking space, the status of the parking space stays 'available', regardless if there are multiple bookings for this parking space.
   
3. Reserved by Admin:
   TODO

4. Arrival of the Driver
    - Once the driver arrives, the parking space status is changed to 'Occupied'.

5. Departure of the Driver
    - Once the driver departs, the parking space status is changed to 'Available', unless Reserved?

6. Blocked by Admin
    - Admin can block parking spaces for reasons like road repair, unexpected problem, which changes the status to 'Blocked', regardless of the current status.
    - No bookings can be made for a parking space once it is blocked.
    - **TODO**: 
        - Affected parking requests are updated
            - pending and disassociate with that parking space.
        - Affected drivers are notified.

7. Unblocked by Admin
    - Goes to 'available'.
