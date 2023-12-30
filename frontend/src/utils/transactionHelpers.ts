export const chooseName = (edges: any) => {
  const { detail, from_account, to_account } = edges
  switch (detail.type) {
    case "transfer":
      if (detail.amount > 0) {
        return `${from_account.edges.user.firstName} ${from_account.edges.user.lastName}`
      } else {
        return `${to_account.edges.user.firstName} ${to_account.edges.user.lastName}`
      }
    case "deposit":
      return `${to_account.edges.user.firstName} ${to_account.edges.user.lastName}`
    case "withdraw":
      return `${from_account.edges.user.firstName} ${from_account.edges.user.lastName}`
    default:
      return "Unknown"
  }
}

export const choosePicture = (edges: any) => {
  const { detail, from_account, to_account } = edges
  switch (detail.type) {
    case "transfer":
      if (detail.amount > 0) {
        return from_account.edges.user.pictureUrl
      } else {
        return to_account.edges.user.pictureUrl
      }
    case "deposit":
      return to_account.edges.user.pictureUrl
    case "withdraw":
      return from_account.edges.user.pictureUrl
    default:
      return "Unknown"
  }
}

export const turnUnique = (transactions: any) => {
  const pictures = transactions.map((transaction: any) => choosePicture(transaction.edges))
  const unique = [...new Set(pictures)]
  return unique
}

export const makeCapitalized = (str: string) => {
  return str.charAt(0).toUpperCase() + str.slice(1)
}
