import { useEffect, useState } from "react";
import { MapContainer, Marker, TileLayer, useMapEvents } from "react-leaflet";
import "leaflet/dist/leaflet.css";

import L from "leaflet";

import type { LeafletMouseEvent } from "leaflet";

const pickupIcon = new L.Icon({
  iconUrl: "https://unpkg.com/leaflet@1.9.4/dist/images/marker-icon.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});

const destinationIcon = new L.Icon({
  iconUrl:
    "https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-red.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
});

export default function CreateRide() {
  const [pickup, setPickup] = useState<[number, number] | null>(null);
  const [destination, setDestination] = useState<[number, number] | null>(null);

  useEffect(() => {
    fetch("http://localhost:2534/api/v001/auth/me", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.Role !== "client") {
          window.location.href = "/me";
        }
      });

    navigator.geolocation.getCurrentPosition((pos) => {
      const newPos: [number, number] = [
        pos.coords.latitude,
        pos.coords.longitude,
      ];

      setPickup(newPos);
      console.log(newPos);
    });
  }, []);

  if (!pickup) {
    return <p>Getting location...</p>;
  }

  function MapClick({
    setDestination,
  }: {
    setDestination: (p: [number, number]) => void;
  }) {
    useMapEvents({
      click(e: LeafletMouseEvent) {
        setDestination([e.latlng.lat, e.latlng.lng]);
      },
    });

    return null;
  }

  async function requestRide() {
    if (!pickup || !destination) alert("Please select pickup and destination");
    else {
      const res = await fetch("http://localhost:2534/api/v001/rides/request", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify({
          start_lat: pickup[0],
          start_lng: pickup[1],
          end_lat: destination[0],
          end_lng: destination[1],
        }),
      });

      if (res.ok) {
        alert("Ride requested successfully");
        const data = await res.json();

        document.getElementById("price")!.textContent =
          `Estimated price: $${Math.floor(data.price).toFixed(2)}`;
        document.getElementById("km")!.textContent =
          `Estimated distance: ${Math.floor(data.distance_km).toFixed(2)} km`;
        console.log(data);
      } else {
        alert("Failed to request ride");
      }
    }
  }

  return (
    <div>
      <h1>Create Ride</h1>

      <MapContainer
        center={pickup}
        zoom={15}
        style={{ height: "400px", width: "100%" }}
      >
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />

        <Marker
          icon={pickupIcon}
          position={pickup}
          draggable={true}
          eventHandlers={{
            dragend: (e: any) => {
              const marker = e.target;
              const pos = marker.getLatLng();

              setPickup([pos.lat, pos.lng]);
            },
          }}
        />

        <MapClick setDestination={setDestination} />

        {destination && (
          <Marker icon={destinationIcon} position={destination} />
        )}
      </MapContainer>

      <p id="price">Estimated price: $0.00</p>
      <p id="km">Estimated distance: 0 km</p>
      <button onClick={requestRide}>Request Ride</button>
    </div>
  );
}
