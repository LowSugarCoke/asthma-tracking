import { questionList } from "./questionData";
import { QuestionItem } from "./questionData";
import React, { useRef, useState, useEffect } from "react";
import * as SplashScreen from "expo-splash-screen";
import {
  ScrollView,
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  Dimensions,
} from "react-native";

const { width } = Dimensions.get("window");

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  slide: {
    width: width,
    justifyContent: "center",
    alignItems: "center",
    padding: 15,
  },
  questionText: {
    fontSize: 22,
    textAlign: "center",
    marginBottom: 20,
  },
  answerButton: {
    marginBottom: 10,
    padding: 10,
    backgroundColor: "#DDDDDD",
    borderRadius: 20,
    width: width - 30,
    alignItems: "center",
  },
  selectedButton: {
    backgroundColor: "#787878",
  },
  dotContainer: {
    flexDirection: "row",
    justifyContent: "center",
    marginTop: 40,
  },
  dot: {
    width: 10,
    height: 10,
    borderRadius: 5,
    backgroundColor: "#CCCCCC",
    marginHorizontal: 5,
  },
  activeDot: {
    backgroundColor: "#2088e8",
  },
});

const QuestionPage = () => {
  const scrollViewRef = useRef<ScrollView>(null);
  const currentIndex = useRef<number>(0);
  const [selected, setSelected] = useState<{ [key: number]: number }>({});

  const handleAnswerPress = (questionIndex: number, answerIndex: number) => {
    if (questionIndex < questionList.length - 1 && scrollViewRef.current) {
      scrollViewRef.current.scrollTo({
        x: width * (questionIndex + 1),
        animated: true,
      });
      currentIndex.current = questionIndex + 1;
      setSelected({
        ...selected,
        [questionIndex]: answerIndex,
      });
    }
  };

  useEffect(() => {
    const hideSplashScreen = async () => {
      await SplashScreen.hideAsync();
    };
    hideSplashScreen();
  }, []);

  return (
    <View style={{ flex: 1 }}>
      <ScrollView
        ref={scrollViewRef}
        style={styles.container}
        horizontal={true}
        pagingEnabled={true}
        showsHorizontalScrollIndicator={false}
        onScroll={(event) => {
          const index = Math.round(
            event.nativeEvent.contentOffset.x /
              event.nativeEvent.layoutMeasurement.width
          );
          currentIndex.current = index;
        }}
        scrollEventThrottle={16}
      >
        {questionList.map((item: QuestionItem, index: number) => (
          <View key={index} style={styles.slide}>
            <Text style={styles.questionText}>{item.question}</Text>
            {item.answer.map((answer: string, answerIndex: number) => (
              <TouchableOpacity
                key={answerIndex}
                style={[
                  styles.answerButton,
                  selected[index] === answerIndex
                    ? styles.selectedButton
                    : null,
                ]}
                onPress={() => handleAnswerPress(index, answerIndex)}
              >
                <Text>{answer}</Text>
              </TouchableOpacity>
            ))}
          </View>
        ))}
      </ScrollView>
      <View style={styles.dotContainer}>
        {questionList.map((_, index) => (
          <View
            key={index}
            style={[
              styles.dot,
              currentIndex.current === index && styles.activeDot,
            ]}
          />
        ))}
      </View>
    </View>
  );
};

export default QuestionPage;
