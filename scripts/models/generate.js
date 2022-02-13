const { GoFileGenerator, FieldType } = require('@asyncapi/modelina');
const Path = require('path');
const doc = require('../../asyncapi.json');
const generator = new GoFileGenerator({
  presets: [
    {
      struct: {
        field({fieldName, content, type }) {
          if (type === 0) {
            return `${content} \`json:"${fieldName}"\``;
          }
          return `${content} \`json:"-"\``;
        },
        additionalContent({renderer, model}) {
          const formattedName = this.nameType(model.$id);
          const listOfRegularProps = Object.entries(model.properties).map(([propName, prop]) => {
            const fieldName = renderer.nameField(propName, prop);
            return `jsonObj["${propName}"] = m.${fieldName}`
          });
          let additionalPropertiesCode = '';
          let unmarshalAdditionalPropertiesCode = '';
          if(model.additionalProperties) {
            additionalPropertiesCode = `for key, element := range m.AdditionalProperties {
  if jsonObj[key] == nil {
    jsonObj[key] = element
  } else {
    fmt.Println("Could not convert additionalProperty Key:", key, " with ", "Element:", element, " to JSON as it already exists.")
  }
}` 

            const additionalPropType = renderer.renderType(model.additionalProperties)
            unmarshalAdditionalPropertiesCode = `else {
  var t ${additionalPropType}
  if c.AdditionalProperties == nil {
    c.AdditionalProperties = make(map[string]${additionalPropType})
  }
  var buf bytes.Buffer
  enc := json.NewEncoder(&buf)
  err := enc.Encode(element)
  if err != nil {
    return err
  }
  UnmarshalJetStreamAccount(buf.Bytes(), &t)
  c.AdditionalProperties[key] = &t
}`
          }

          const regularPropsUnmarshal = Object.entries(model.properties).map(([propName, prop]) => {
            const fieldName = renderer.nameField(propName, prop);
            let fieldType = renderer.renderType(prop);
            return `if key == "${propName}" {
              c.${fieldName} = element.(${fieldType})
            }`
          });
          
          return `
func (m *${formattedName}) MarshalJSON() ([]byte, error) {
  jsonObj := make(map[string]interface{})
  ${listOfRegularProps.join('\n')}

  ${additionalPropertiesCode}
  return json.Marshal(jsonObj)
}

func Unmarshal${formattedName}(body []byte, c *${formattedName}) error {
  jsonObj := make(map[string]interface{})
  err := json.Unmarshal(body, &jsonObj)
  if err != nil {
    return err
  }

  for key, element := range jsonObj {
    fmt.Println("Converting Key:", key, " with ", "Element:", element)
    ${regularPropsUnmarshal.join(' else ')} 
  }
  return nil
}
          
          `
        }
      }
    }
  ]
});
const outputDir = Path.resolve(__dirname, '../../internal/models')
generator.generateToFiles(doc, "" + outputDir, { packageName: 'api' });