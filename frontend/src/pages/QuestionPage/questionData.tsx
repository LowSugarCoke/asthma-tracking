export interface QuestionItem {
  question: string;
  answer: string[];
}

export const questionList: QuestionItem[] = [
  {
    question: "Did you experience any asthma-related coughing today?",
    answer: [
      "a) No, I'm feeling great today~~~",
      "b) I've had some mild coughing.",
      "c) I've been coughing due to asthma for most of the day and finding it difficult to breathe.",
      "d) I'm currently coughing and experiencing difficulty breathing as a result of asthma.",
    ],
  },
  {
    question:
      "Have you experienced wheezing during the day as a result of asthma?",
    answer: [
      "a) No, I'm feeling great today~~~",
      "b) Occasionally, I do feel some wheezing.",
      "c) I experience wheezing for most of the day and find it difficult to breathe.",
      "d) I constantly struggle with breathing and find it unpleasant",
    ],
  },
  {
    question:
      "Does your asthma impact your ability to run, exercise, and participate in sports?",
    answer: [
      "a) No, I'm feeling good today~~",
      "b) I experience symptoms when I feel excited.",
      "c) I'm unable to engage in activities with my friends.",
      "d) I constantly struggle with breathing and it restricts me from doing what I want to do.",
    ],
  },
  {
    question: "Do you experience nighttime awakenings due to your asthma?",
    answer: [
      "a) No, I had a restful sleep last night.",
      "b) Yes, occasionally.",
      "c) Yes, most of the time.",
      "d) Yes, consistently throughout the night.",
    ],
  },
  {
    question:
      "How many doses of your reliever medication have you taken in the past 24 hours?",
    answer: [
      "a) I haven't taken any doses today.",
      "b) I have taken two doses.",
      "c) I have taken six doses.",
      "d) I have taken more than six doses, but it hasn't provided relief.",
    ],
  },
  {
    question:
      "How frequently did you experience wheezing during the day due to your asthma?",
    answer: [
      "a) No, I'm feeling good today~",
      "b) Twice.",
      "c) Six times.",
      "d) More than six times.",
    ],
  },
];
