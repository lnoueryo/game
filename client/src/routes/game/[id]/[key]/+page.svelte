<script lang="ts">
    import { page } from '$app/stores'
	import { io } from 'socket.io-client';
	import type { PageData } from './$types';
    import type { player, card } from '@prisma/client';
    export let data: PageData

    const socket = io();
    let user = data.user;
    let table = $page.data.table;
    let start = $page.data.table.start;
    let cards = $page.data.table.cards

	socket.on('table', (_table) => {
		if(String(_table.key) != $page.params.key) return;
		table = _table;
		cards = _table.cards;
        console.log(cards)
        if(!start && table.start) {
            start = true;
        }
	});

    const filterPlayers = (players: player[]) => {
        return players.filter(player => player.id != user.id)
    }

    const startGame = async() => {
        const res = await fetch('http://127.0.0.1:5173/api/table', {
            method: 'PUT',
            body: JSON.stringify({key: $page.params.key, players: table.players})
        });
        const newTable = await res.json()
        socket.emit('table', newTable);
    }

    const handCards = (cards: card[], player: player) => {
        return cards.filter(card => card.playerId == player.id)
    }
    if(table.players.length == table.limit) {
        startGame()
    }
    socket.emit('table', table);
</script>

<div class="container">
    <div>
        {#if filterPlayers(table.players).length == 0}
            <div>プレイヤーの参加を待っています。</div>
        {:else}
            <div class="d-flex">
                {#each filterPlayers(table.players) as player}
                <div class="card mb-3" style="max-width: 540px;">
                    <div class="row g-0">
                        <div class="col-md-4">
                            
                        </div>
                        <div class="col-md-8">
                            <div class="card-body">
                            <h5 class="card-title">{player.username}</h5>
                            <p class="card-text">level</p>
                            <p class="card-text"><small class="text-muted">Last updated 3 mins ago</small></p>
                            </div>
                        </div>
                        {#each handCards(cards, player) as card}
                            <div>{card.name}{card.type}</div>
                        {/each}
                    </div>
                </div>
                {/each}
            </div>
        {/if}
    </div>
    <div>
        <div class="card mb-3" style="max-width: 540px;">
            <div class="row g-0">
                <div class="col-md-4">
                    <img>
                </div>
                <div class="col-md-8">
                    <div class="card-body">
                    <h5 class="card-title">{user.username}</h5>
                    <p class="card-text">level</p>
                    <p class="card-text"><small class="text-muted">Last updated 3 mins ago</small></p>
                    </div>
                </div>
                {#each handCards(cards, user) as card}
                    <div>{card.name}{card.type}</div>
                {/each}
            </div>
        </div>
    </div>
</div>