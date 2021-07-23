<template lang="pug">
q-page
  .row.justify-center.q-pt-xl
    .col-6
      q-select(
        hide-dropdown-icon,
        outlined,
        v-model="model",
        use-input,
        hide-selected,
        fill-input,
        input-debounce="0",
        placeholder="Search ...",
        :options="options",
        @filter="filterFn",
        @filter-abort="abortFilterFn"
      )
        template(v-slot:no-option)
          q-item
            q-item-section.text-grey
              | No results
</template>

<script>
import { useRoute, useRouter } from "vue-router"
import { ref, watch } from "vue"
const stringOptions = ["Google", "Facebook", "Twitter", "Apple", "Oracle"]

export default {
  setup() {
    const options = ref(stringOptions)
    const model = ref(null)

    let router = useRouter()

    watch(model, (model, prevModel) => {
      router.push({ path: "/lineage" })
    })
    return {
      model,
      options,
      filterFn(val, update, abort) {
        // call abort() at any time if you can't retrieve data somehow

        if (val.length < 2) {
          abort()
          return
        }

        setTimeout(() => {
          update(() => {
            if (val === "") {
              options.value = stringOptions
            } else {
              const needle = val.toLowerCase()
              options.value = stringOptions.filter(
                (v) => v.toLowerCase().indexOf(needle) > -1
              )
            }
          })
        }, 1500)
      },
      abortFilterFn() {
        // console.log('delayed filter aborted')
      },
    }
  },
}
</script>
