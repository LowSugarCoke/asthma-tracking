import React from "react";
import { Button } from "react-native";
import { useNavigation } from "@react-navigation/native";

const LoginPage = () => {
  const navigation = useNavigation();

  return (
    <div>
      <h1> Login Page</h1>;
      <Button
        title="Go to Patient Info"
        onPress={() => navigation.navigate("patientinfo")}
      />
    </div>
  );
};

export default LoginPage;
