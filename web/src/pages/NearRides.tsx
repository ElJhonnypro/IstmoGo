import { useEffect, useState } from "react";
import { MapContainer, TileLayer, Marker } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import L from "leaflet";

type Ride = {
  id: string;
  start_lat: number;
  start_lng: number;
  end_lat: number;
  end_lng: number;

  price: number;
};
import type { UserInfo } from "../assets/types/models";
const rideIcon = new L.Icon({
  iconUrl:
    "https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-green.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});

export default function NearRides() {
  const [rides, setRides] = useState<Ride[]>([]);
  const [location, setLocation] = useState<[number, number] | null>(null);
  const [User, setUser] = useState<UserInfo | null>(null);

  useEffect(() => {
    navigator.geolocation.getCurrentPosition((pos) => {
      const lat = pos.coords.latitude;
      const lng = pos.coords.longitude;

      setLocation([lat, lng]);

      fetch("http://localhost:2534/api/v001/rides/getNearRides", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify({
          lat,
          lng,
        }),
      })
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          setRides(data.rides);
        });
    });

    fetch("http://localhost:2534/api/v001/auth/me", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log("User info", data);
        setUser(data);
        console.log("User info state", User);
      });
  }, []);

  async function acceptRide(id: string) {
    const res = await fetch("http://localhost:2534/api/v001/rides/accept", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify({ ride_id: id }),
    });

    const data = await res.json();

    console.log(data);

    if (res.ok) {
      alert("Ride accepted");
      window.location.href = "/rides/rider/current";
    }
  }

  if (!location) {
    return <p>Getting location...</p>;
  }

  return (
    <div>
      <h1>Near Rides</h1>

      <MapContainer
        center={location}
        zoom={14}
        style={{ height: "400px", width: "100%" }}
      >
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />

        {rides.map((ride) => (
          <Marker
            key={ride.id}
            icon={rideIcon}
            position={[ride.start_lat, ride.start_lng]}
            eventHandlers={{
              click: () => {
                acceptRide(ride.id);
              },
            }}
          />
        ))}
      </MapContainer>

      <h2>Available rides</h2>

      <ul>
        {rides.map((ride) => (
          <li key={ride.id}>
            <p>
              From: {ride.start_lat}, {ride.start_lng}
            </p>
            <p>
              To: {ride.end_lat}, {ride.end_lng}
            </p>
            <p>Price: ${ride.price}</p>

            <button onClick={() => acceptRide(ride.id)}>Accept Ride</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
