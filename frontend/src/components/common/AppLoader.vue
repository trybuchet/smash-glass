<template>
    <div class="app-loader">
        <svg xmlns="http://www.w3.org/2000/svg" version="1.1">
            <defs>
                <filter id="fancy-goo">
                    <feGaussianBlur in="SourceGraphic" stdDeviation="10" result="blur" />
                    <feColorMatrix in="blur" mode="matrix" values="1 0 0 0 0  0 1 0 0 0  0 0 1 0 0  0 0 0 19 -9" result="goo" />
                </filter>
            </defs>
        </svg>

        <div class="container top-con">
            <div class="blob-container">
                <div class="blob"></div>
            </div>
            <div class="blob-container">
                <div class="blob"></div>
            </div>
            <div class="blob-container">
                <div class="blob"></div>
            </div>
            <div class="blob-container">
                <div class="blob"></div>
            </div>
            <div class="blob-container">
                <div class="blob"></div>
            </div>
            <div class="blob-container">
                <div class="blob"></div>
            </div>
        </div>
        <div class="container bottom-con"></div>
		<h2>Please Wait...</h2>
    </div>
</template>
<script lang="ts">

export default {
    name: 'AppLoader',
    components: {
    },
    data() {
        return {
        };
    },
    computed: {
        
    },
    methods: {
        
    },
    mounted() {
    },
}
</script>
<style lang="scss" scoped>
.app-loader {
	text-align: center;
}
.bottom-con {
	z-index: -1;
}

.top-con {
	filter: url('#fancy-goo');
	// animation-name: rotate
	animation-iteration-count: infinite;
	animation-duration: 2s;
	animation-timing-function: ease-in-out;
}

.container {
	width: 300px;
	height: 100px;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
}
		
.element {
	width: 100%;
	position: absolute;
	bottom: 0;
	height: 130px;
	background: white;
}

.blob-container {
	width: 100%;
	height: 30%;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);

	&:nth-child(1) {
		transform: translate(-50%, -50%) rotate(-95deg);
		.blob {
			animation-delay: 0ms;
        }
    }
	&:nth-child(2) {
		transform: translate(-50%, -50%) rotate(-85deg);
		.blob {
			animation-delay: 350ms;
        }
    }
	&:nth-child(3) {
		transform: translate(-50%, -50%) rotate(-88deg);
		.blob {
			animation-delay: 800ms;
        }
    }
		
	&:nth-child(4) {
		transform: translate(-50%, -50%) rotate(-92deg);
		.blob {
			animation-delay: 1200ms;
        }
    }

	&:nth-child(5) {
		transform: translate(-50%, -50%) rotate(-120deg);
		.blob {
			animation-delay: 850ms;
        }
    }

	&:nth-child(6) {
		transform: translate(-50%, -50%) rotate(-60deg);
		.blob {
			animation-delay: 500ms;
        }
    }
		
	.blob {
		left: 50%;
		top: 50%;
		width: 65px;
		height: 65px;
		background: white;
		position: absolute;
		transform: translate(-50%, -50%);
		bottom: 95px;
		border-radius: 99em;
		animation-name: move;
		animation-iteration-count: infinite;
		animation-duration: 1600ms;
		animation-timing-function: ease-in-out;
		background: linear-gradient(to top right, #f32c3c, #FF802B);
    }

}

	
@keyframes move {
	0% {
		left: 50%;
		top: 50%;
		width: 75px;
		height: 65px;
    }

	30% {
		left: 50%;
		top: 50%;
		width: 65px;
		height: 85px;
    }

	70% {
		left: 90%;
		top: 50%;
		width: 25px;
		height: 25px;
    }

	90% {
		left: 55%;
		top: 50%;
		width: 1px;
		height: 1px;
    }
	100% {
		left: 50%;
		top: 50%;
		width: 65px;
		height: 75px;
    }
}
</style>