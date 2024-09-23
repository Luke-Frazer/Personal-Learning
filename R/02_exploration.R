library(tidyverse)
library(palmerpenguins)
library(rmarkdown)
options(dplyr.width = Inf)

data("penguins")
penguins

#       rows        columns
# c() makes the columns a vector
penguins[1:5, c("species", "island")]

# when you just specify one column, you dont need c()
penguins[1:5, "species"]

# returns all row data for variable ${var}
penguins$flipper_length_mm

# For indexing, you need to specify which rows you wish to return
# leaving col blank is a wildcard
penguins[penguins$island == "Dream", ]

penguins %>%
  group_by(species) %>%
  summarise(averageFlipperLen = mean(flipper_length_mm, na.rm = TRUE))

# This counts the number of each species on each island
penguins %>%
  group_by(species, island) %>%
  summarise(numPenguinsPerIsland = n())

# Could also do this, equivalient v
penguins %>%
  group_by(species, island) %>%
  count()

# puts year first, followed by the rest (everything else)
penguins %>%
  select(year, everything())

# makes column 'size_group' that checks if each penguin has body mass > average and if so, set to large, otherwise small # nolint
penguins %>%
  mutate(size_group = if_else(body_mass_g >= mean(body_mass_g, na.rm = TRUE), "large", "small")) %>% # nolint
  select(size_group)

# gives only rows where the year isnt 2008/9
penguins %>% filter(!(year %in% c(2008, 2009))) %>% arrange(desc(year))

ggplot(penguins) +
  geom_smooth(
    mapping = aes(x = body_mass_g, y = flipper_length_mm, color = species)
  )