const fs = require('fs');
const yaml = require('js-yaml');

const YAML_PATH = "./CA.yaml"

function saveData(key, value) {
  try {
    let yamlContent = {};
    
    // 如果文件存在，读取现有内容
    if (fs.existsSync(YAML_PATH)) {
        const fileContent = fs.readFileSync(YAML_PATH, 'utf8');
        yamlContent = yaml.load(fileContent) || {};
    }

    // 更新或添加新的键值对
    yamlContent[key] = value;

    // 将更新后的内容写回文件
    const newYamlContent = yaml.dump(yamlContent);
    fs.writeFileSync(YAML_PATH, newYamlContent, 'utf8');

    console.log(`Successfully updated key "${key}" with value:`, value);
    } catch (error) {
        console.error('Error updating YAML file:', error);
        throw error;
  }
}

function getYamlValue(key) {
  try {
      // 确保文件存在
      if (!fs.existsSync(YAML_PATH)) {
          throw new Error('YAML file does not exist');
      }

      // 读取并解析 YAML 文件
      const fileContent = fs.readFileSync(YAML_PATH, 'utf8');
      const yamlContent = yaml.load(fileContent);

      // 检查键是否存在
      if (!(key in yamlContent)) {
          throw new Error(`Key "${key}" not found in YAML file`);
      }

      return yamlContent[key];
  } catch (error) {
      console.error('Error reading YAML file:', error);
      throw error;
  }
}

module.exports = {
    saveData,
    getYamlValue
}