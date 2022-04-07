const FormData = (form:HTMLFormElement) => {
    const inputs = form.querySelectorAll("input")
    let values: {[prop: string]: string} = {}
    inputs.forEach(input=>{
        values[input.name] = input.value
    })
    return values
}
export default FormData