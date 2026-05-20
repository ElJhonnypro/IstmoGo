import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { useEffect, useState } from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";

import Login from "../screens/Login";
import Register from "../screens/Register";
import CreateRide from "../screens/CreateRide";
import NearRides from "../screens/NearRides";
import MapViewLocation from "../screens/MapViewLocation";

const Stack = createNativeStackNavigator();

export default function AppNavigator() {
  const [user, setUser] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchUser = async () => {
      const token = await AsyncStorage.getItem("token");

      if (!token) {
        setLoading(false);
        return;
      }

      fetch("http://localhost:2534/api/v001/auth/me", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
        .then((res) => res.json())
        .then((data) => {
          setUser(data);
          setLoading(false);
        })
        .catch(() => {
          setLoading(false);
        });
    };
    fetchUser();
  }, []);

  if (loading) return null;

  return (
    <Stack.Navigator screenOptions={{ headerShown: false }}>
      {!user ? (
        <>
          <Stack.Screen name="Login" component={Login} />
        </>
      ) : user.role === "client" ? (
        <>
          <Stack.Screen name="CreateRide" component={CreateRide} />
        </>
      ) : (
        <>
          <Stack.Screen name="NearRides" component={NearRides} />
        </>
      )}
      <Stack.Screen name="Register" component={Register} />
      <Stack.Screen name="MapViewLocation" component={MapViewLocation} />
    </Stack.Navigator>
  );
}
