import type { player, table } from "@prisma/client"

export const createCards = (key: string) => {
    const types = ['ハート', 'ダイヤ', 'スペード', 'ジャック']
    const cards = [
      {id: crypto.randomUUID(), name: 'JOKER', type: 'ジョーカー', active: true, tableId: key},
      {id: crypto.randomUUID(), name: 'JOKER', type: 'ジョーカー', active: true, tableId: key},
    ]
    types.forEach((type) => {
      for (let i = 1; i < 14; i++) {
        cards.push({id: crypto.randomUUID(), name: String(i), type: type, active: true, tableId: key})
      }
    })
    const shuffledCards = shuffle(cards)
    return shuffledCards
}

const shuffle = ([...array]) => {
    for (let i = array.length - 1; i >= 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
}