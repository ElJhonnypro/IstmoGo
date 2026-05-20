import { useState, useEffect } from "react";

import type { UserInfo } from "../assets/types/models";

export default function Me() {
  const [userInfo, setUserInfo] = useState<UserInfo | null>(null);

  function getUserInfo() {
    const token = localStorage.getItem("token");
    fetch("http://localhost:2534/api/v001/auth/me", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log("User info", data);
        setUserInfo(data);
        console.log("User info state", userInfo);
      });
  }

  useEffect(() => {
    getUserInfo();
  }, []);

  return (
    <div>
      <h1>Me</h1>
      <p>Usuario: {userInfo?.Name}</p>
    </div>
  );
}
