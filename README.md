# terraform-provider-epic

## Ideas
1. Customizable Templates for Names
Allow users to define patterns or templates for names that include placeholders which your provider can fill. For example, users might specify a pattern like "${character}-${descriptor}-${unique_suffix}", where:

${character} could be a random character from a selected media.
${descriptor} could be an adjective or a role associated with that character.
${unique_suffix} could be a random numeric or alphanumeric string for uniqueness.
2. Richer Descriptions Using Quotes
Besides using a single quote, allow the construction of more complex descriptions. You could combine multiple quotes or add additional context or commentary programmatically.

3. Dynamic Media Selection
Allow users to specify which media source (e.g., "Star Wars", "Lord of the Rings", "Game of Thrones") to use for generating names, either globally or per resource. This could be configured through provider attributes or directly in the resource definition.

4. Tagging with Attributes
Generate tags automatically based on media, character traits, or themes. For instance, names generated from "Game of Thrones" could include tags like {"media": "Game of Thrones", "house": "Stark"} based on the character's affiliation.

5. Include Contextual Information
For each generated name, provide a context object that might include the character’s full name, title, affiliations, and a brief description or history. This could be useful for documentation or managing resources related to specific characters or themes.

6. Interactive Configuration
Develop a CLI tool or a setup wizard that helps users configure the provider interactively, selecting their preferred media sources, setting up name patterns, and other preferences.

9. Advanced Text Processing
Incorporate natural language processing (NLP) techniques to generate names and descriptions that are not only random but also contextually relevant and linguistically varied.

11. Extendibility
Design the provider so that users can extend it with their own data files or plugins. This would allow the community to contribute additional media types or character sets.
