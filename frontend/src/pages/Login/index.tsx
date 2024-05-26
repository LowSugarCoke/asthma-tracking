import React from "react";
import { Button } from "react-native";
import { useNavigation } from "@react-navigation/native";

const LoginPage = () => {
  const navigation = useNavigation();

  return (
    <div>
      <h1> Welcome Asthma Tracking</h1>
      <Button
        title="Go to Question Page"
        onPress={() => navigation.navigate("questionpage")}
      />
    </div>
  );
};

export default LoginPage;
