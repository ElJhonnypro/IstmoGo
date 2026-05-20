import { useEffect, useState } from "react";
import { MapContainer, Marker, TileLayer } from "react-leaflet";
import type { RideInfoData } from "../assets/types/models";

import "leaflet/dist/leaflet.css";
import L from "leaflet";

// ICONS
import UMIconPNG from "../assets/UMIcon.png";
import PLIconPNG from "../assets/PLIcon.png";

const UMIcon = new L.Icon({
  iconUrl: UMIconPNG,
  iconSize: [41, 21],
  iconAnchor: [12, 41],
});

const DIcon = new L.Icon({
  iconUrl:
    "https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-red.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});

const PLIcon = new L.Icon({
  iconUrl: PLIconPNG,
  iconSize: [41, 41],
  iconAnchor: [12, 41],
});

export default function CurrentRideUside() {
  const [ULocation, setULocation] = useState<[number, number] | null>(null);
  const [PLocation, setPLocation] = useState<[number, number] | null>(null);
  const [DLocation, setDLocation] = useState<[number, number] | null>(null);
  const [rideId, setRideId] = useState<string | null>(null);
  function move(dLat: number, dLng: number) {
    if (!ULocation) return;

    const newPos: [number, number] = [ULocation[0] + dLat, ULocation[1] + dLng];

    setULocation(newPos);
  }

  // =========================
  // DRIVER LOCATION (REALTIME)
  // =========================
  useEffect(() => {
    const watchId = navigator.geolocation.watchPosition((pos) => {
      const lat = pos.coords.latitude;
      const lng = pos.coords.longitude;

      console.log("New position:", lat, lng);

      setULocation([lat, lng]);
    });

    return () => {
      navigator.geolocation.clearWatch(watchId);
    };
  }, []);

  // =========================
  // GET RIDE INFO (INTERVAL)
  // =========================
  useEffect(() => {
    fetch("http://localhost:2534/api/v001/rides/getMyRide", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    })
      .then((res) => res.json())
      .then((data: RideInfoData) => {
        if (data.rides && data.rides.length > 0) {
          const ride = data.rides[0];

          console.log("Ride info:", ride);

          setRideId(ride.id);

          setPLocation([ride.start_lat, ride.start_lng]);
          setDLocation([ride.end_lat, ride.end_lng]);
        }
      })
      .catch((err) => console.log("error fetching ride:", err)); // cada 5 segundos
  }, []);

  // =========================
  // LOADING STATE
  // =========================
  if (!ULocation) {
    return <p>Getting location...</p>;
  }

  return (
    <div>
      <h1>Current Ride</h1>

      <MapContainer
        center={ULocation}
        zoom={15}
        style={{ height: "400px", width: "100%" }}
      >
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />

        {/* DRIVER (moving smoothly) */}
        <Marker icon={UMIcon} position={ULocation} />

        {/* PICKUP */}
        {PLocation && <Marker icon={PLIcon} position={PLocation} />}

        {/* DESTINATION */}
        {DLocation && <Marker icon={DIcon} position={DLocation} />}
      </MapContainer>

      <div>
        <h1>Finish Ride?</h1>
        <button
          onClick={() => {
            console.log("Finishing ride with ID:", rideId);
            fetch("http://localhost:2534/api/v001/rides/finish", {
              method: "POST",
              body: JSON.stringify({
                ride_id: rideId,
              }),
              headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("token")}`,
              },
            })
              .then((res) => res.json())
              .then((data) => {
                console.log("Ride finished:", data);
                alert("Ride finished!");
              })
              .catch((err) => console.log("error finishing ride:", err));
          }}
        >
          Finish Ride
        </button>
        <button onClick={() => move(0.001, 0)}>↑</button>
        <button onClick={() => move(-0.001, 0)}>↓</button>
        <button onClick={() => move(0, 0.001)}>→</button>
        <button onClick={() => move(0, -0.001)}>←</button>
      </div>
    </div>
  );
}
