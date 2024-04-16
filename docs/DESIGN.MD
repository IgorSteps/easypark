# Design

This document describes the states, rules and assumptions of entities within the EasyPark system. It is a guide for admins and developers in managing and understanding how parking resources are allocated and maintained.

## States

### Parking Lot

- When a parking lot is created, all its parking spaces in it are of `available` status.
- When a parking lot is deleted, all its parking spaces and parking requests referencing those parking spaces are deleted.

### Parking Request

States a parking request can be in: `pending`, `approved`, `rejected`, `active` and `approved`.

- **Pending**:
  - All new parking requests are created a 'pending' status.
  - If the admin 'reserves' or 'blocks' a space assigned to parking request, it will change request's status to 'pending' and disassociate with that parking space in the system.
- **Approval Conditions**:
  - A parking request can only be approved if:
    - It is currently in 'pending' status.
    - Its time slot does not overlap with existing parking requests for the designated parking space.
    - The requested start time is in the future.
    - The parking space is in the requested parking lot and is neither 'blocked' nor 'reserved'.
- **Rejection**:
  - A parking request will be rejected if the admin cannot find a parking space that suites driver's request, such as:
    - There are no available spaces that match driver's desired time.
- **Active**:
  - A parking request is 'active' if:
    - It received an arrival notification
    - It was previously 'approved'
- **Completed**:
  - A parking request is 'completed' if:
    - It received a departure notification
    - It was previously 'active'

### Parking Space

States a parking space can be in: `available`, `occupied`, `reserved`, `blocked`.

- **Initial State**:
  - Upon creation, all parking spaces are set to 'available'.
- **Assignment**:
  - A parking space remains 'available' even after being assigned to a driver’s parking request, allowing for multiple requests to be considered(as long as they don't overlap).
- **Admin Reservation**:
  - Admin can reserve spaces for specific purposes, such as events. Reserved spaces cannot be assigned to parking requests.
  - A reservation starts immediately and only ends when the admin changes the status of the parking space.
- **Unreserving**:
  - When a space is unreserved by the admin, its status immediately returns to 'available'.
- **Driver Arrival**:
  - When a driver arrives and confirms their location, the status of the assigned parking space changes to 'occupied'.
- **Driver Departure**:
  - Once the driver departs, the status reverts to 'available'.
- **Admin Blocking**:
  - Spaces can be blocked by the admin for road repairs or emergencies. No bookings can be made for blocked spaces.
  - A block starts immediately and only ends when the admin changes the status of the parking space.
- **Unblocking**:
  - When a space is unblocked by the admin, its status immediately returns to 'available'.

### Alerts

**Scheduling**:

Checks for alerts are scheduled using the [Scheduler](../internal/drivers/scheduler/scheduler.go) to run at constant intervals of 5 minutes.

**Late arrival alert**:

All parking requests:
    - in past must have status 'completed'.
    - current but that have received an arrival notification have status 'active'.
  
So when checking for late arrivals, we look at parking requests that are before 1 hour in the past and of status 'approved'.

## Assumptions

- If the admin 'blocks' or 'reserves' a parking space with existing parking requests assigned to it, it is their responsibility to notify the driver of that change and reassign a parking space to their parking request.
- If the admin 'rejects' a parking request, it is their responsibility to notify the driver of the reason, so that the driver can make another parking request that can get approved.
- If multiple parking requests compete for the same time slot, it falls onto the admin to choose which request gets approved and which gets rejected.
- If the admin receives an alert that location mismatch has happened on arrival, it is assumed the admin will respond to the alert to prevent operational disruptions using our messaging system to guide the driver to the correct parking space.
- If the admin receives an alert that an arrival notification has not been received within one hour from the booked arrival time, it is assumed they will act accordingly to prevent operational disruptions, such as:
  - Driver has arrived but forgot to send a notification:
    - Admin updates space status to 'occupied'
  - Driver hasn't arrived:
    - No action needed, as the space status is 'available'.
- If the admin receives an alert that a departure notification has not been received within one hour after the booked departure time, it is assumed they will act accordingly to prevent operational disruptions, such as:
  - Driver has left, but forgot to send a notification:
    - Admin updates space status to 'available'
  - Driver hasn't left:
    - Admin contacts the driver to tell him to leave and to send a departure notification.
    - Admin reassigns the spaces of the affected parking requests.

## Future work

- Introducing a "Parking Restriction" feature will allow for better management of parking space reservations and blockages over time. This would support multiple, time-based restrictions per space but is not planned for the initial release (MVP).
