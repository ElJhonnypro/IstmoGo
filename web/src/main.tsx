import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import Register from "./pages/Register";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Me from "./pages/Me";
import CreateRide from "./pages/CreateRide";
import GetMerides from "./pages/getmerides";
import NearRides from "./pages/NearRides";
import CurrentRideUside from "./pages/CurrentRideUside";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/Me" element={<Me />} />
        <Route path="/create-ride" element={<CreateRide />} />
        <Route path="/me/ride" element={<GetMerides />} />
        <Route path="/uber/near-rides" element={<NearRides />} />
        <Route path="/uber/currentride" element={<CurrentRideUside />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>,
);
