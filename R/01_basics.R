library(tidyverse)
library(palmerpenguins)

summary <- penguins %>%
  select(bill_depth_mm) %>%
  summarise(medBillDepth = median(bill_depth_mm, na.rm = TRUE))

summary

# combines species and islands that are alike
group <- penguins %>%
  group_by(species, island)

group

mutation <- penguins %>%
  mutate(bill_length_cm = bill_length_mm / 10)

mutation

females <- penguins %>%
  filter(species == "Adelie" & island == "Torgersen")

females

ggplot(penguins) +
  geom_point(
    mapping = aes(x = bill_length_mm, y = bill_depth_mm, color = species)
  )
