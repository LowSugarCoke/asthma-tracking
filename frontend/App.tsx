import React from "react";
import { PaperProvider } from "react-native-paper";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import QuestionPage from "./src/pages/QuestionPage";
import LoginPage from "./src/pages/Login";

const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <PaperProvider>
      <NavigationContainer>
        <Stack.Navigator
          screenOptions={{
            headerShown: false,
          }}
        >
          <Stack.Screen name="login" component={LoginPage} />
          <Stack.Screen name="questionpage" component={QuestionPage} />
        </Stack.Navigator>
      </NavigationContainer>
    </PaperProvider>
  );
}
