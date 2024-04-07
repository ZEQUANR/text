const generateFormModel = () => {
  return {
    number: "",
    name: "",
    contentType: "",
    filterType: "",
    createdTime: [],
    status: "",
  }
}

generateFormModel().number = 10

console.log(generateFormModel())
