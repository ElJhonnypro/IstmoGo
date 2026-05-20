import { useEffect } from "react"

export default function GetMerides() {
    useEffect(() => {
        fetch("http://localhost:2534/api/v001/rides/getMyRide", {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
            },
        })
        .then(res => res.json())
        .then(data => {
            console.log(data)
            document.getElementById("price")!.textContent = `Estimated price: $${Math.floor(data.ride[1].price).toFixed(2)}`
            document.getElementById("km")!.textContent = `Estimated distance: ${Math.floor(data.ride[1].distance_km).toFixed(2)} km`
        })
    }, [])
    return (
        <div>
            <h1>My Ride</h1>
            <p id="price">Estimated price: </p>
            <p id="km">Estimated distance: </p>
        </div>
    )
}