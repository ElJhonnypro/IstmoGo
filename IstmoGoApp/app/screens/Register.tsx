import { View, Text, Image, Pressable } from "react-native";
import { NavigationProp, useNavigation } from "@react-navigation/native";
import { styles as LStyles } from "../Styles/Components/LoginStyles";
import { styles as BStyles } from "../Styles/General/Backgrounds";
import radius from "../Styles/Components/radius";
import Inputs from "../Styles/Components/Inputs";
import JInputText from "../components/Inputs";
import colors from "../Styles/Components/colors";
import * as Location from "expo-location";
import { useEffect, useState } from "react";
import { RootStackParamList } from "../types/navigation";

export default function Login({ route }: any) {
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();
  const dataLocation = route.params?.dataLocation || null;
  const [name, setName] = useState("");
  const [phone, setPhone] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [location, setLocation] = useState({
    latitude: 0,
    longitude: 0,
  });
  const [birthdate, setBirthdate] = useState("");
  const [role, setRole] = useState("client");

  // uber fields
  const [rid, setRID] = useState("");
  const [carPlate, setCarPlate] = useState("");
  const [carModel, setCarModel] = useState("");
  const [carColor, setCarColor] = useState("");
  const [carPhoto, setCarPhoto] = useState<File | null>(null);

  const [ridPhoto, setRidPhoto] = useState<File | null>(null);

  async function handleRegister() {
    let carId = null;

    // =========================
    // 1 Register Car if uber
    // =========================
    if (role === "uber") {
      const carForm = new FormData();

      carForm.append("plate", carPlate);
      carForm.append("model", carModel);
      carForm.append("color", carColor);

      if (carPhoto) {
        carForm.append("photo", carPhoto);
      }

      const carRes = await fetch(
        "http://localhost:2534/api/v001/car/register",
        {
          method: "POST",
          body: carForm,
        },
      );

      const carData = await carRes.json();

      console.log("car response", carData);

      carId = carData.carId;
    }

    // =========================
    // 2 Register User
    // =========================

    const form = new FormData();

    form.append("name", name);
    form.append("phone", phone);
    form.append("role", role);
    form.append("email", email);
    form.append("password", password);
    form.append("locationLat", location.latitude.toString());
    form.append("locationLong", location.longitude.toString());
    form.append("birthdate", birthdate);

    if (role === "uber") {
      form.append("rid", rid);

      if (ridPhoto) {
        form.append("photo", ridPhoto);
      }

      form.append("carId", carId);
      console.log("carId", carId);
    }

    const res = await fetch("http://localhost:2534/api/v001/auth/register", {
      method: "POST",
      body: form,
    });

    const data = await res.json();
    console.log(password);
    console.log("user response", data);
  }

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();

      if (status !== "granted") {
        alert("Permiso de ubicación denegado");
        return;
      }

      const loc = await Location.getCurrentPositionAsync({});

      setLocation({
        latitude: loc.coords.latitude,
        longitude: loc.coords.longitude,
      });
    })();
  }, []);

  function OpenMap() {
    alert("Selecciona tu ubicación en el mapa");
    navigation.navigate("MapViewLocation", {
      goBack: "Register",
      data: [],
    });
  }

  return (
    <View style={[BStyles.main, LStyles.container]}>
      <View
        style={[
          BStyles.black,
          {
            borderRadius: radius.Normal,
            padding: 5,
            margin: 5,
            minHeight: 574,
            minWidth: 346,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "flex-start",
          },
        ]}
      >
        <Image
          style={[
            LStyles.logo,
            {
              flex: 0,
            },
          ]}
          source={require("../../assets/images/LogoHorizontal.png")}
        />

        <View style={LStyles.Form}>
          <JInputText
            placeholder="Nombre de usuario"
            value={name}
            onChangeText={setName}
          />
          <JInputText
            placeholder="Teléfono"
            value={phone}
            onChangeText={setPhone}
            keyboardType="phone-pad"
          />
          <JInputText
            placeholder="Correo electrónico"
            value={email}
            onChangeText={setEmail}
            keyboardType="email-address"
            autoCapitalize="none"
            autoCorrect={false}
          />
          <JInputText
            placeholder="Contraseña"
            value={password}
            onChangeText={setPassword}
            secureTextEntry
          />

          <Pressable
            style={({ pressed }) => [
              Inputs.Button,
              {
                backgroundColor: pressed
                  ? colors.primary + "CC"
                  : colors.primary,

                color: pressed ? colors.Third + "CC" : colors.Third,
              },
            ]}
            onPress={OpenMap}
          >
            <Text>{dataLocation ? "Change location" : "Select location"}</Text>
          </Pressable>
          <Text>{dataLocation}</Text>
          <Pressable
            style={({ pressed }) => [
              Inputs.Button,
              {
                backgroundColor: pressed
                  ? colors.secondary + "CC"
                  : colors.secondary,

                color: pressed ? colors.Third + "CC" : colors.Third,
              },
            ]}
            onPress={handleRegister}
          >
            <Text style={Inputs.ButtonText}>Registrarse</Text>
          </Pressable>

          <Pressable onPress={() => navigation.navigate("Login" as never)}>
            <Text style={Inputs.link}>¿Ya tienes una cuenta? Logueate</Text>
          </Pressable>
        </View>
      </View>
    </View>
  );
}
