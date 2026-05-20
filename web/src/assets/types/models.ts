export type UserInfo = {
    Birthdate
: 
string

CarID
: 
string
CreatedAt
: 
string
Email
: 
string
ID
: 
string
Location
: 
string
Name
: 
string
Password
: 
string
Phone
: 
string
RID
: 
string
RIDPhoto
: 
string
Role
: 
string
}



export type RideInfoData = {
    message: string
    rides: Array<{
        id: string
        client_id: string
        uber_id: string
        start_lat: number
        start_lng: number
        end_lat: number
        end_lng: number
        distance_km: number
        price: number
        status: string
        requested_at: string
    }>
}

