<template>
    <div class="box">
        <div class="control" style="margin-bottom: 10px;">
            <textarea class="textarea is-primary " id="message" type="text " placeholder="Message here " value=" "></textarea>
        </div>
        <div class="level">
            <div class="level-left ">
                <div class=" control level-item has-icons-left">

                    <div class="select ">
                        <select>
                            <option>*All</option>
                            <option>*Yourself</option>
                        </select>
                        <div class="icon is-small is-left">
                            <i class="fas fa-envelope"></i>
                        </div>
                    </div>

                </div>
            </div>

            <br>
            <div class="level-right">
                <div class="level-item">
                    <a class="button is-primary " onclick="send(); ">Send Message</a>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'SendBox',
    props: {
        uri: String,
        name: String
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
    margin: 40px 0 0;
}
ul {
    list-style-type: none;
    padding: 0;
}
li {
    display: inline-block;
    margin: 0 10px;
}
a {
    color: #42b983;
}
</style>
